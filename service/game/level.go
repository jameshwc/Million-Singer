package game

import (
	"strconv"

	"github.com/jameshwc/Million-Singer/model"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"gorm.io/gorm"
)

func GetLevel(param string) (*model.Level, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return nil, C.ErrLevelIDNotNumber
	}
	level, err := model.GetLevel(id)
	if err == gorm.ErrRecordNotFound {
		return nil, C.ErrLevelNotFound
	} else if err != nil {
		return nil, C.ErrDatabase
	}
	return level, nil
}
