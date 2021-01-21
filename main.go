package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"knowledgeBase/src/controllers"
)

func main() {
	//r := gin.New()
	//r.GET("/user", func(c *gin.Context) {
	//	user := UserModel.NewUseModel().Mutate(UserModel.WithUserID(001)).
	//		Mutate(UserModel.WithUserName("Hendrix"))
	//	c.JSON(200, user)
	//})
	//r.GET("/log", func(c *gin.Context) {
	//	user :=UserLog.NewUserLog().Mutate(UserLog.WithUserLogIp("1.1.1.1"))
	//	c.JSON(200, user)
	//})
	//r.Run(":9091")

	goft.Ignite().Mount("v1",controllers.NewUserController()).Launch()
}
