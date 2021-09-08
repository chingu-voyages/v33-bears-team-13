package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
	// "os"
	"fmt"
	"log"
	owm "github.com/briandowns/openweathermap"
	"net/http"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

// var apiKey = os.Getenv("API_KEY")

func APIHandler(ctx *gin.Context) {
	city := ctx.Param("city")

	w, err := owm.NewCurrent("C", "en", apiKey) // fahrenheit (imperial) with Russian output //
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName(city)
	fmt.Println(w)
	ctx.JSON(http.StatusOK, w)
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	server := gin.Default()
    server.Use(CORS())
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

    server.GET("/weather", func(ctx *gin.Context) {
	    ctx.JSON(200, gin.H{
			"message": "Let's get weather",
		})
    })



	server.GET("/weather/:city", func(ctx *gin.Context) {
        APIHandler(ctx)
    })

	server.Run(":8080")
}
