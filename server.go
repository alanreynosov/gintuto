package main

import (
	"github.com/gin-gonic/gin"
	"github.com/alanreynosov/gintuto/controller"
	"github.com/alanreynosov/gintuto/service"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main(){
	server := gin.Default()
	server.GET("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})


	server.POST("/videos", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	server.Run(":8080")
}

