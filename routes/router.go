package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ylinyang/vue-demo-go/api/v1"
	"github.com/ylinyang/vue-demo-go/middleware"
	"github.com/ylinyang/vue-demo-go/model"
	"github.com/ylinyang/vue-demo-go/utils"
	"log"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	s := gin.Default()

	auth := s.Group("/api/v1")
	auth.POST("user/add", v1.AddUser)
	auth.POST("login", v1.Login)

	auth.Use(middleware.JwtToken())
	auth.GET("test", model.Test)

	s.Group("/api/v1")
	if err := s.Run(utils.HttpPort); err != nil {
		log.Panicln(err)
	}
}
