package subjects

import (
	"errors"
	"modules/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SubjectsRoutes(router *gin.RouterGroup) {
	router.GET("/", getAllSubjects)
}


func getAllSubjects(c *gin.Context) {
	subjectModels, modelCount, err := FindManySubjects("20", "0")
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("subjects", errors.New("Something went wrong")))
		return
	}
	/* serializer := SubjectsSerializer{c, subjectModels} */
	c.JSON(http.StatusOK, gin.H{"subjects": /* serializer.Response() */subjectModels, "subjectsCount": modelCount})
}