package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ylinyang/vue-demo-go/middleware"
	"github.com/ylinyang/vue-demo-go/model"
	"github.com/ylinyang/vue-demo-go/utils"
	"net/http"
)

func Login(context *gin.Context) {
	var data model.User
	context.ShouldBindJSON(&data)

	c := model.CheckLogin(data.Username, data.Password)
	if c == utils.SUCCESS {
		token := middleware.SetToken(data.Username)
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": utils.GetErrMsg(c),
			"token":   token,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  c,
		"message": utils.GetErrMsg(c),
	})
}
