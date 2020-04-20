package controller

import (
    "github.com/ft-ironman/golang-gin-poc/entity"
    "github.com/ft-ironman/golang-gin-poc/service"
    "github.com/gin-gonic/gin"
)

type VideoController interface {
    FindAll() []entity.Video
    Save(context *gin.Context) entity.Video
}

type controller struct {
    service service.VideoService
}

func New(service service.VideoService) VideoController {
    return controller{
        service: service,
    }
}

func (c *controller) FindAll() []entity.Video {
    return service.FindAll()
}

func (c *controller) Save(ctx gin.Context) entity.Video {
    var Video entity.Video
    ctx.BindJSON(&Video)
    c.service.Save(Video)
    return Video
}