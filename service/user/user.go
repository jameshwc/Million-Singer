package user

import (
	"github.com/astaxie/beego/validation"
	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/pkg/token"
)

type user struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50); MinSize(6);"`
}

type checkUser struct {
	Username string `valid:"MaxSize(50)"`
	Email    string `valid:"MaxSize(100)"`
}

type registerUser struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50); MinSize(6);"`
	Email    string `valid:"Required; MaxSize(100)"`
}

func AuthUser(username, password string) (string, error) {
	// TODO: email
	checkUser := user{Username: username, Password: password}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&checkUser)
	if !ok {
		return "", C.ErrUserLoginFormat
	}
	u, err := model.AuthUser(username, password)
	if err != nil {
		return "", C.ErrUserLoginAuthentication
	}
	userToken, err := token.Generate(username, password)
	if err != nil {
		return "", C.ErrUserLoginJwtTokenGeneration
	}
	if err := u.UpdateLoginStatus(); err != nil {
		return "", C.ErrUserLoginUpdateUserStatus
	}
	return userToken, nil
}

func ValidateUser(username string, email string) error {
	if username == "" && email == "" {
		return C.ErrUserCheckParamIncorrect
	}
	valid := validation.Validation{}
	var u checkUser
	ok, _ := valid.Valid(&u)
	if !ok {
		return C.ErrUserCheckFormat
	}
	if username != "" && model.IsUserNameDuplicate(username) {
		return C.ErrUserCheckNameConflict
	}
	if email != "" && model.IsUserEmailDuplicate(email) {
		return C.ErrUserCheckEmailConflict
	}
	return nil
}

func Register(username, email, password string) error {
	valid := validation.Validation{}
	ok, _ := valid.Valid(&registerUser{
		Username: username, Email: email, Password: password,
	})
	if !ok {
		return C.ErrUserRegisterFormat
	}
	if model.IsUserEmailDuplicate(email) {
		return C.ErrUserRegisterEmailConflict
	}
	if model.IsUserNameDuplicate(username) {
		return C.ErrUserRegisterNameConflict
	}
	id, err := model.AddUser(username, email, password)
	if err != nil {
		log.Error("error when register a user: ", err)
		return C.ErrUserRegisterFailServerError
	}
	log.Infof("user %d has registered with username %s", id, username)
	return nil
}
