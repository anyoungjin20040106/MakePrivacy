package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.File("index.html")
	})
	r.POST("/privacy", func(ctx *gin.Context) {
		serviceName := ctx.PostForm("service_name")
		companyName := ctx.PostForm("company_name")
		email := ctx.PostForm("email")
		collectedData := ctx.PostFormArray("collected_data")
		data := gin.H{
			"ServiceName":   serviceName,
			"CompanyName":   companyName,
			"Email":         email,
			"CollectedData": collectedData,
		}
		ctx.HTML(http.StatusOK, "result.html", data)

	})
	r.NoRoute(func(ctx *gin.Context) {
		ctx.File("404.html")
	})
	r.Run(":1234")
}
