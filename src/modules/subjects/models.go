package subjects

import (
	"fmt"
	"strconv"
	"modules/common"
)

type Subject struct {
  	ID uint `json:”id”`
	Slug string `json:slug`
	Subject string `json:”subject"`
}

func FindManySubjects(limit, offset string) ([]Subject, int, error) {
	db := common.GetDB()
	var models []Subject
	var count int

	offset_int, err := strconv.Atoi(offset)
	fmt.Println(err)
	if err != nil {
		offset_int = 0
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = 20
	}

	tx := db.Begin()
	db.Model(&models).Count(&count)
	db.Offset(offset_int).Limit(limit_int).Find(&models)

	err = tx.Commit().Error
	return models, count, err
}