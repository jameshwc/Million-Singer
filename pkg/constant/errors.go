package constant

import "errors"

var (
	// Login Error
	ErrUserLoginFormat             = errors.New("incorrect format of username or password")
	ErrUserLoginAuthentication     = errors.New("incorrect username or password")
	ErrUserLoginJwtTokenGeneration = errors.New("fail to generate jwt token")
	ErrUserLoginUpdateUserStatus   = errors.New("fail to update user's login status")

	ErrUserCheckParamIncorrect = errors.New("no username or email in params")
	ErrUserCheckFormat         = errors.New("incorrect format of username or email")
	ErrUserCheckNameConflict   = errors.New("username conflict")
	ErrUserCheckEmailConflict  = errors.New("email conflict")

	ErrUserRegisterFailServerError = errors.New("fail to create the user")
	ErrUserRegisterFormat          = errors.New("incorrect format of username or email or password")
	ErrUserRegisterNameConflict    = errors.New("username conflict")
	ErrUserRegisterEmailConflict   = errors.New("email conflict")
)