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
	limit := c.Query("limit")
	offset := c.Query("offset")
	subjectModels, modelCount, err := FindManySubjects(limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("subjects", errors.New("Something went wrong")))
		return
	}
	/* serializer := SubjectsSerializer{c, subjectModels} */
	c.JSON(http.StatusOK, gin.H{"subjects": /* serializer.Response() */subjectModels, "subjectsCount": modelCount})
}