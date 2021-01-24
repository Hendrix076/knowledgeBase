package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"knowledgeBase/src/data/Getter"
	"knowledgeBase/src/middlewares"
	"knowledgeBase/src/models/UserModel"
)

func UserInfoList(c *gin.Context) {
	req := &UserModel.ResUserInfo{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(400, gin.H{"result": err})
	}
	userlist, err := Getter.UserGetter.GetUserList(req.Users)
	if err != nil {
		c.JSON(400, gin.H{"result": err})
	}
	c.JSON(200, gin.H{"result": userlist})
}
func UserAccessToken(c *gin.Context) {
	 phone, pwd := c.PostForm("phone"), c.PostForm("password")
	if phone == "" && pwd == "" {
		c.JSON(400, gin.H{"result": "用户和密码不能为空"})
	}
	user,err:=Getter.UserGetter.FindByphone(phone)
	if err !=nil {
		fmt.Println(err)
	}
	if user.Password == pwd {
		token := middlewares.GenerateToken(phone)
		c.JSON(200, gin.H{"token": token})
	}
}
