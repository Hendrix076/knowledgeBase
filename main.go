package main

import (
	"github.com/gin-gonic/gin"
	"knowledgeBase/src/common"
	"knowledgeBase/src/controllers"
	"knowledgeBase/src/middlewares"
)

func main() {
	common.InitDB()
	r := gin.New()
	r.Use(middlewares.ErrorHandler())
	r.Handle("POST", "/userinfo", controllers.UserInfoList)
	r.Handle("POST", "/access_token", controllers.UserAccessToken)
	r.Run(":8080")
}
