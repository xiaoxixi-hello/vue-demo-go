package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ylinyang/vue-demo-go/model"
	"github.com/ylinyang/vue-demo-go/utils"
	"net/http"
	"strconv"
)

var code int

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)

	if code == utils.SUCCESS {
		model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": utils.GetErrMsg(code),
	})
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	fmt.Println(pageSize, pageNum)
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = 1
	}
	users := model.GetUsers(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": utils.SUCCESS,
		"data":   users,
	})
}
