package message

import "errors"

var (
	NameRequired       = errors.New("field name is required")
	UserNameRequired   = errors.New("field username is required")
	EmailRequired      = errors.New("field email is required")
	PasswordRequired   = errors.New("field password is required")
	TokenRequired      = errors.New("field token is required")
	FullNameRequired   = errors.New("field name is required")
	IdRequired         = errors.New("field id is required")
	AuthHeaderRequired = errors.New("authorization header is required")
)

var (
	UserNotFound          = errors.New("user not found")
	WrongUserNamePassword = errors.New("wrong username or password")
	DuplicateData         = errors.New("data already exists")
	InternalServerError   = errors.New("internal server error")
)

var (
	HeaderFormatInvalid = errors.New("invalid authorization header format")
)

var (
	AccessDenied       = "you dont have access to this menu, access denied"
	InvalidToken       = "invalid token"
	TokenExpired       = "token is expired"
	ModuleNameRequired = "module name is required"
	FilePathRequired   = "file path is required"
)
