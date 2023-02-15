package main

import (
	"go-rest-api/controllers"
	"go-rest-api/models"
	"log"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
	docs "go-rest-api/docs"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	models.ConnectDatabase()

	docs.SwaggerInfo.BasePath = "/api/v1"
   v1 := r.Group("/api/v1")
   {
			v1.GET("/ping", controllers.Ping)
      v1.GET("/songs", controllers.FindSongs)
			v1.POST("/songs", controllers.CreateSong)
			v1.GET("/songs/:id", controllers.FindSong)
			v1.PUT("/songs/:id", controllers.UpdateSong)
			v1.DELETE("/songs/:id", controllers.DeleteSong)
   }
   r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Println("Starting server.")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}