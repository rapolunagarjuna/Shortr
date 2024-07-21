package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/rapolunagarjuna/Shortr/handlers"
)

func main() {
	r := gin.Default()

	urls := make(map[string]string)
	sHandler := handlers.ShortenHandler{Urls: urls}

	r.GET("/", func(c *gin.Context) {
		temp := template.Must(template.ParseFiles("templates/HomePage.html"))
		temp.Execute(c.Writer, nil)
	})

	r.GET("/login", func(c *gin.Context) {
		temp := template.Must(template.ParseFiles("templates/LoginPage.html"))
		temp.Execute(c.Writer, nil)
	})

	r.GET("/profile", func(c *gin.Context) {
		temp := template.Must(template.ParseFiles("templates/ProfilePage.html"))
		temp.Execute(c.Writer, nil)
	})

	r.GET("/ping", handlers.PingHandler)

	r.POST("/shorten", sHandler.AddNewLink)
	r.GET("/:alias", sHandler.Redirect)
	r.GET("/qrcode", handlers.QRCodeHandler)

	r.Run()
}
