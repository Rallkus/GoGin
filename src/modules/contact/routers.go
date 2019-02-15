package contact

import (
	"github.com/gin-gonic/gin"
)

func SubjectsRoutes(router *gin.RouterGroup) {
	router.GET("/", sendEmail)
}

func sendEmail(c *gin.Context){
	return 
}