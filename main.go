package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	Ads := os.Getenv("Ads")
	if Ads == "" {
		Ads = "Default Ad Text"
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Ads": Ads,
		})
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
			"Ads":           Ads,
		}
		ctx.HTML(http.StatusOK, "result.html", data)

	})
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "404.html", gin.H{
			"Ads": Ads,
		})
	})
	r.Run(":8000")
}
