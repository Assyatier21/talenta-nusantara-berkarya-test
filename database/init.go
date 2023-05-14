package database

import (
	"tnb-technical-test/models"
	"tnb-technical-test/utils/constant"
)

func Init() *[]models.ProgrammingLanguage {
	return &constant.GlobalLanguage
}
