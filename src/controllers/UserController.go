package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"knowledgeBase/src/data/Getter"
	"knowledgeBase/src/middlewares"
	"knowledgeBase/src/models/UserModel"
)
//获取用户列表
func UserInfoList(c *gin.Context) {
	req := &UserModel.ResUserInfo{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		panic(errors.New("请求参数绑定失败"))
	}
	userlist, err := Getter.UserGetter.GetUserList(req.Users)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"result": userlist})
}
//生成token
func UserAccessToken(c *gin.Context) {
	 phone, pwd := c.PostForm("phone"), c.PostForm("password")
	if phone == "" || pwd == "" {
		panic(errors.New("用户和密码不能为空"))

	}
	user,err:=Getter.UserGetter.FindByphone(phone)
	if err !=nil {
		panic(err)
	}
	if user.Password == pwd {
		token := middlewares.GenerateToken(phone)
		c.JSON(200, gin.H{"token": token})
	}
}
