package main

import (
	"github.com/gin-gonic/gin"
	"knowledgeBase/src/common"
	"knowledgeBase/src/controllers"
)

func main() {
	common.InitDB()
	r := gin.New()
	r.Handle("POST", "/userinfo", controllers.UserInfoList)
	r.Handle("POST", "/access_token", controllers.UserAccessToken)
	r.Run(":8080")
}
