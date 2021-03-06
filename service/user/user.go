package user

import (
	"database/sql"

	"github.com/astaxie/beego/validation"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/pkg/token"
	"github.com/jameshwc/Million-Singer/repo"
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

func (src *Service) Auth(username, password string) (string, error) {
	// TODO: email
	checkUser := user{Username: username, Password: password}
	valid := validation.Validation{}
	ok, _ := valid.Valid(&checkUser)
	if !ok {
		return "", C.ErrUserLoginFormat
	}
	u, err := repo.User.Auth(username, password)
	if err == sql.ErrNoRows {
		return "", C.ErrUserLoginAuthentication
	} else if err != nil {
		return "", C.ErrDatabase
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

func (src *Service) Validate(username string, email string) error {
	if username == "" && email == "" {
		return C.ErrUserCheckParamIncorrect
	}
	valid := validation.Validation{}
	var u checkUser
	ok, _ := valid.Valid(&u)
	if !ok {
		return C.ErrUserCheckFormat
	}
	if username != "" && repo.User.IsNameDuplicate(username) {
		return C.ErrUserCheckNameConflict
	}
	if email != "" && repo.User.IsEmailDuplicate(email) {
		return C.ErrUserCheckEmailConflict
	}
	return nil
}

func (src *Service) Register(username, email, password string) error {
	valid := validation.Validation{}
	ok, _ := valid.Valid(&registerUser{
		Username: username, Email: email, Password: password,
	})
	if !ok {
		return C.ErrUserRegisterFormat
	}
	if repo.User.IsEmailDuplicate(email) {
		return C.ErrUserRegisterEmailConflict
	}
	if repo.User.IsNameDuplicate(username) {
		return C.ErrUserRegisterNameConflict
	}
	id, err := repo.User.Add(username, email, password)
	if err != nil {
		log.Error("error when register a user: ", err)
		return C.ErrUserRegisterFailServerError
	}
	log.Infof("user %d has registered with username %s", id, username)
	return nil
}
