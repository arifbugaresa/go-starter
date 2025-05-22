package user

import (
	"encoding/json"
	"github.com/arifbugaresa/go-starter/middlewares"
	"github.com/arifbugaresa/go-starter/utils/common"
	"github.com/arifbugaresa/go-starter/utils/common/message"
	"github.com/arifbugaresa/go-starter/utils/constant/enum"
	"github.com/arifbugaresa/go-starter/utils/response"
	"github.com/arifbugaresa/go-starter/utils/session"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"time"
)

type Service interface {
	Login(ctx *gin.Context, dataBody LoginRequest) (response LoginResponse, err error)
	Logout(ctx *gin.Context) (err error)
	RegisterUser(ctx *gin.Context, dataBody RegisterRequest) (response SignUpResponse, err error)

	GetUserProfile(ctx *gin.Context) (response GetProfileResponse, err error)
	UpdateUserProfile(ctx *gin.Context, dataBody UpdateProfileRequest) (err error)
}

type UserService struct {
	repo  *UserRepository
	redis *redis.Client
}

func NewService(repo *UserRepository, redis *redis.Client) *UserService {
	return &UserService{
		repo:  repo,
		redis: redis,
	}
}

func (s *UserService) RegisterUser(ctx *gin.Context, dataBody RegisterRequest) (response SignUpResponse, err error) {
	var (
		timeNow = time.Now()
	)

	passwordHashed, err := common.HashPassword(dataBody.Password)
	if err != nil {
		return
	}

	roleUser, err := s.repo.GetRoleByRoleName(ctx, enum.User)
	if err != nil {
		return
	}

	err = common.WrapperTx(ctx, s.repo.db, func(tx *sqlx.Tx) (err error) {
		modelUser := RegisterModel{
			FullName: dataBody.FullName,
			UserName: dataBody.UserName,
			Email:    dataBody.Email,
			Password: passwordHashed,
			Photo:    dataBody.Photo,
			RoleId:   roleUser.RoleId,
			DefaultTable: common.DefaultTable{
				CreatedAt: common.DefaultFormatDate(timeNow),
				UpdatedAt: common.DefaultFormatDate(timeNow),
			},
		}

		err = s.repo.RegisterUserTx(ctx, modelUser, tx)
		if err != nil {
			return
		}

		return
	})
	if err != nil {
		return
	}

	return
}

func (s *UserService) Login(ctx *gin.Context, dataBody LoginRequest) (response LoginResponse, err error) {
	user, err := s.CheckUserAndPassword(ctx, dataBody)
	if err != nil {
		return
	}

	// generate token
	claims := middlewares.Claims{
		Role:     user.Role,
		RoleId:   user.RoleId,
		FullName: user.FullName,
		UserName: user.UserName,
		Email:    user.Email,
		Photo:    user.Photo,
	}

	jwtToken, err := claims.GenerateJwtToken()
	if err != nil {
		return
	}

	err = s.SetSessionToRedis(ctx, user, jwtToken)
	if err != nil {
		return
	}

	response.Token = jwtToken

	return
}

func (s *UserService) CheckUserAndPassword(ctx *gin.Context, dataBody LoginRequest) (user LoginModel, err error) {
	user, err = s.repo.GetUserByEmailOrUsername(ctx, LoginModel{UserName: dataBody.UserName})
	if err != nil {
		return
	}

	if user.Id == 0 {
		err = message.UserNotFound
		return
	}

	// check password user
	matches := common.CheckPassword(user.Password, dataBody.Password)
	if !matches {
		err = message.WrongUserNamePassword
		return
	}

	return
}

func (s *UserService) SetSessionToRedis(ctx *gin.Context, user LoginModel, jwtToken string) (err error) {
	redisSession := session.RedisData{
		Id:       user.Id,
		Photo:    user.Photo,
		FullName: user.FullName,
		UserName: user.UserName,
		Role:     user.Role,
		RoleId:   user.RoleId,
		Email:    user.Email,
	}

	jsonBytes, err := json.Marshal(redisSession)
	if err != nil {
		return
	}

	s.redis.Set(ctx, jwtToken, string(jsonBytes), 1*time.Hour)

	return
}

func (s *UserService) Logout(ctx *gin.Context) (err error) {
	token, err := middlewares.GetJwtTokenFromHeader(ctx)
	if err != nil {
		return
	}

	// delete session from redis
	s.redis.Del(ctx, token)

	return
}

func (s *UserService) GetUserProfile(ctx *gin.Context) (response GetProfileResponse, err error) {
	auth, err := middlewares.GetSession(ctx)
	if err != nil {
		return
	}

	user, err := s.repo.GetUserByEmailOrUsername(ctx, LoginModel{UserName: auth.UserName})
	if err != nil {
		return
	}

	response = GetProfileResponse{
		Id:       user.Id,
		UserName: user.UserName,
		FullName: user.FullName,
		Email:    user.Email,
		Photo:    user.Photo,
		Role:     user.Role,
		RoleId:   user.RoleId,
	}

	return
}

func (s *UserService) UpdateUserProfile(ctx *gin.Context, dataBody UpdateProfileRequest) (err error) {
	auth, err := middlewares.GetSession(ctx)
	if err != nil {
		response.GenerateErrorResponse(ctx, err.Error())
	}

	data := UpdateProfileModel{
		Id:           auth.Id,
		FullName:     dataBody.FullName,
		Photo:        &dataBody.Photo,
		DefaultTable: common.DefaultTable{}.GetDefaultTable(ctx),
	}

	err = s.repo.UpdateUserById(ctx, data)
	if err != nil {
		return
	}

	return
}
