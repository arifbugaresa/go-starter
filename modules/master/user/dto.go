package user

import (
	"github.com/arifbugaresa/go-starter/utils/common"
	"github.com/arifbugaresa/go-starter/utils/common/message"
)

type (
	RegisterRequest struct {
		FullName string `json:"full_name"`
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Photo    string `json:"photo"`
	}

	RegisterModel struct {
		Email    string `db:"email"`
		UserName string `db:"username"`
		Password string `db:"password"`
		Photo    string `db:"photo" goqu:"omitempty"`
		FullName string `db:"full_name"`
		RoleId   int64  `db:"role_id"`
		common.DefaultTable
	}

	SignUpResponse struct{}
)

type (
	GetRoleModel struct {
		RoleId   int64  `db:"role_id"`
		RoleName string `db:"role_name"`
		common.DefaultTable
	}
)

func (r RegisterRequest) ValidateRegisterRequest() error {
	if r.FullName == "" {
		return message.NameRequired
	}

	if r.UserName == "" {
		return message.UserNameRequired
	}

	if r.Email == "" {
		return message.UserNameRequired
	}

	if r.Password == "" {
		return message.PasswordRequired
	}

	return nil
}

type (
	LoginRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Token string `json:"token"`
	}

	LoginModel struct {
		Id       int64   `db:"id"`
		RoleId   int64   `db:"role_id"`
		UserName string  `db:"username"`
		Email    string  `db:"email"`
		FullName string  `db:"full_name"`
		Password string  `db:"password"`
		Role     string  `db:"role"`
		Photo    *string `db:"photo" goqu:"omitnil"`
	}
)

func (r LoginRequest) ValidateLoginRequest() error {
	if r.UserName == "" {
		return message.UserNameRequired
	}

	if r.Password == "" {
		return message.PasswordRequired
	}

	return nil
}

type (
	LogoutRequest struct {
		Token string `json:"token"`
	}
)

func (r LogoutRequest) ValidateLogoutRequest() (err error) {
	if r.Token == "" {
		return message.TokenRequired
	}

	return
}

type (
	GetProfileRequest struct {
		Id int64 `json:"id"`
	}

	GetProfileResponse struct {
		Id       int64   `json:"id"`
		UserName string  `json:"username"`
		FullName string  `json:"full_name"`
		Email    string  `json:"email"`
		Photo    *string `json:"photo"`
		Role     string  `json:"role"`
		RoleId   int64   `json:"role_id"`
	}

	GetProfileModel struct {
		Id       int64   `db:"id"`
		FullName string  `db:"full_name"`
		Email    string  `db:"email"`
		Photo    *string `db:"photo"`
		Role     string  `db:"role"`
	}

	UpdateProfileRequest struct {
		FullName string `json:"full_name"`
		Photo    string `json:"photo"`
	}

	UpdateProfileModel struct {
		Id       int64   `db:"id" goqu:"skipupdate"`
		FullName string  `db:"full_name"`
		Photo    *string `db:"photo" goqu:"omitnil"`
		common.DefaultTable
	}
)

func (r GetProfileRequest) ValidateGetDetailProfileRequest() (err error) {
	if r.Id == 0 {
		return message.IdRequired
	}

	return
}

func (r UpdateProfileRequest) ValidateUpdateProfileRequest() (err error) {
	if r.FullName == "" {
		return message.FullNameRequired
	}

	return
}

type (
	SignUpBaristaRequest struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Photo    string `json:"photo"`
	}

	SignUpBaristaResponse struct{}
)

func (r SignUpBaristaRequest) ValidateSignUpBaristaRequest() error {
	if r.Email == "" {
		return message.UserNameRequired
	}

	if r.Password == "" {
		return message.PasswordRequired
	}

	return nil
}

type (
	DeleteBaristaRequest struct {
		Id int64 `json:"id"`
	}

	DeleteModel struct {
		Id int64 `db:"id"`
	}

	DeleteBaristaResponse struct{}
)

func (r DeleteBaristaRequest) ValidateDeleteBaristaRequest() error {
	if r.Id == 0 {
		return message.IdRequired
	}

	return nil
}

type (
	GetListBaristaRequest struct {
		common.DefaultListRequest
	}

	GetListBaristaResponse struct {
		Id       string  `json:"id"`
		Email    string  `json:"email"`
		FullName string  `json:"full_name"`
		Photo    *string `json:"photo"`
	}

	GetListBaristaModel struct {
		Id       string  `db:"id"`
		Email    string  `db:"email"`
		FullName string  `db:"full_name"`
		Photo    *string `db:"photo" goqu:"omitempty"`
	}
)
