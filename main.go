package main

import (
	"github.com/ylinyang/vue-demo-go/model"
	"github.com/ylinyang/vue-demo-go/routes"
)

func main() {
	// 引用数据库
	model.InitDb()
	routes.InitRouter()
}
