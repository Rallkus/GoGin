package teachers

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	"github.com/wangzitian0/golang-gin-starter-kit/common"
	"subjects"
	"strconv"
)
type TeachersModel struct {
	gorm.Model
	Name        string `gorm:"unique_index"`
	Surname     string `gorm:"unique_index"`
	Sex			string `gorm:"size:2048"`
	Age    		uint
	DNI			string `gorm:"size:2048"`
	Subject		[]SubjectModel `gorm:"ForeignKey:Subject"`
}