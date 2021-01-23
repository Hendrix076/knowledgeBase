package controllers

import (

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"knowledgeBase/src/models/UserModel"
	"knowledgeBase/src/service"
)

type UserController struct {
	userservice *service.UserService
}

type LoginResponse struct {
	Token string `json:"token"` // jwt令牌
}

func NewUserController() *UserController {
	return &UserController{}
}

func (this *UserController) Name() string {
	return "UserController"
}

func (this *UserController) UserInfoList(c *gin.Context) goft.Json {
	  req:=&UserModel.ResUserInfo{}
	goft.Error(c.ShouldBindJSON(req))
	userinfolist,err := this.userservice.Userinfo(req.Users)
fmt.Println(err)

	return gin.H{"result": userinfolist}

}

//func (this *UserController) AccessToken(c *gin.Context) {
//	var (
//		req  = new(UserModel.UserModel)
//		resp = new(LoginResponse)
//		err  error
//	)
//
//	// 请求json绑定
//	phone := c.PostForm("phone")
//	password := c.PostForm("password")
//	if len(phone) == 0 || len(password) == 0 {
//		errors.New("账号和密码不能为空")
//		return
//	}
//
//	req.Phone = phone
//	req.Password = password
//
//	u := &UserModel.UserModel{
//		Phone:    req.Phone,
//		Password: req.Password,
//	}
//	// 密码校验
//	user, err := middleware.LoginCheck(u)
//	if err != nil {
//		return
//	}
//
//	token, err := middleware.GenUserToken(user)
//	if err != nil {
//		errors.New("获取usertoken失败")
//		return
//	}
//	resp.Token = token
//	c.JSON(200,resp.Token)
//
//}

func (this *UserController) Build(goft *goft.Goft) {

	goft.Handle("POST", "/userinfo", this.UserInfoList)
}
