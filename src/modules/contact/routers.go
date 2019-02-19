package contact

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func ContactRoutes(router *gin.RouterGroup) {
	router.POST("/", sendEmail)
}

func sendEmail(c *gin.Context) {
	fmt.Println(c.Params)
	from := mail.NewEmail("Example User", "serhuegi@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "serhuegi@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return
}
