package main

import (
	"github.com/gin-gonic/gin"
	"knowledgeBase/src/models/UserModel"

	UserLog "knowledgeBase/src/models/UserlogModel"
)

func main() {
	r := gin.New()
	r.GET("/user", func(c *gin.Context) {
		user := UserModel.NewUseModel().Mutate(UserModel.WithUserID(001)).
			Mutate(UserModel.WithUserName("Hendrix"))
		c.JSON(200, user)
	})
	r.GET("/log", func(c *gin.Context) {
		user :=UserLog.NewUserLog(UserLog.WithUserLogID(23323), UserLog.WithUserLogIp("1.1.1.1.1"))
		c.JSON(200, user)
	})
	r.Run(":9091")
}
