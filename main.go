package main

import (
	"app/models"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("public/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{"User": user, "Errors": getErrors(err, user)})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{"Pass": true, "User": user})
	})
	router.Run()
}

func getErrors(err error, obj any) map[string]string {
	messages := getMessages(obj)
	errors := map[string]string{}
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = messages[e.Field()]
	}
	return errors
}

func getMessages(obj any) map[string]string {
	t := reflect.TypeOf(obj)
	messages := map[string]string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		messages[field.Name] = field.Tag.Get("msg")
	}
	return messages
}
