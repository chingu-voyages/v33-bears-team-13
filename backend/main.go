package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/chingu-voyages/v33-bears-team-13/backend/controller"
	"github.com/chingu-voyages/v33-bears-team-13/backend/repository"
	"github.com/chingu-voyages/v33-bears-team-13/backend/service"

	owm "github.com/briandowns/openweathermap"
)

var ct, _ = context.WithTimeout(context.Background(), 50*time.Second)


var (
	summaryRepository, _ = repository.Open(ct) 
	summaryService    service.SummaryService       = service.New(summaryRepository)
	summaryController controller.SummaryController = controller.New(summaryService)
	
)




func APIHandler(ctx *gin.Context) *owm.CurrentWeatherData {
	city := ctx.Param("city")
	// apiKey := os.Getenv("API_KEY")
	apiKey := os.Getenv("API_KEY")
	w, err := owm.NewCurrent("C", "en", apiKey) // fahrenheit (imperial) with Russian output //
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName(city)
	fmt.Println(w)


	// ctx.JSON(http.StatusOK, w)

	return w
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

    err := godotenv.Load(".env")
    
    if err != nil {
        log.Fatal("Error loading .env file")
    }


	r := gin.Default()
    r.Use(CORS())



	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

    r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello,World!")
    })


	

	r.GET("/summaries", func(ctx *gin.Context) {
		// var list, err = summaryController.FindAll(ct) ([]string, error)
		r_, err := summaryController.FindAll(ct);
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(r_)
		}
		ctx.JSON(200, r_)
	})



	// r.GET("/summaries", list)
	



	r.POST("/summaries", func(ctx *gin.Context) {

		err := summaryController.Save(ctx, ct)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
		}

	})


	r.GET("/weather/:city", func(ctx *gin.Context) {

		
        w := APIHandler(ctx)


        jsonBytes, err := json.Marshal(w)
		if err != nil {
			panic(err)
		}
		// data := make(map[string]interface{})

		var data *owm.CurrentWeatherData

		var err2 = json.Unmarshal(jsonBytes, &data)
		if err2 != nil {
			log.Fatal(err2)
		}
		var res = "Current temperature in " + data.Name + " is " + strconv.FormatFloat(data.Main.Temp, 'f', 1, 64) + "°C .................Conditions are currently " + data.Weather[0].Description;
		
		ctx.JSON(200, res)
    })


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
