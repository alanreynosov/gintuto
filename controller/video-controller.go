package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/alanreynosov/gintuto/entity"
	"github.com/alanreynosov/gintuto/service"
)

type VideoCotroller interface {
	FindAll() []entity.Video
	Save(c *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoCotroller {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}