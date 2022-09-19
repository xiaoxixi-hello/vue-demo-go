package model

import (
	"github.com/gin-gonic/gin"
	"github.com/ylinyang/vue-demo-go/utils"
	"net/http"
)

func Test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  utils.SUCCESS,
		"message": "加油",
	})
}
