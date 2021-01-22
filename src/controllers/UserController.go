package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"knowledgeBase/src/middleware"
	"knowledgeBase/src/models/UserModel"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (this *UserController) Name() string {
	return "UserController"
}

func (this *UserController) Create(c *gin.Context){
	req:=UserModel.ResUserInfo{}
	err:= c.ShouldBind(req)
	if err !=nil {
		return
	}
}

type LoginResponse struct {
	Token string `json:"token"` // jwt令牌
}

func AccessToken(c *gin.Context) {
	var (
		req  = new(UserModel.UserModel)
		resp = new(LoginResponse)
		err  error
	)

	// 请求json绑定
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	if len(phone) == 0 || len(password) == 0 {
		errors.New("账号和密码不能为空")
		return
	}

	req.Phone = phone
	req.Password = password

	u := &UserModel.UserModel{
		Phone:    req.Phone,
		Password: req.Password,
	}
	// 密码校验
	user, err := middleware.LoginCheck(u)
	if err != nil {
		return
	}

	token, err := middleware.GenUserToken(user)
	if err != nil {
	errors.New("获取usertoken失败")
		return
	}
	resp.Token = token

}

func (this *UserController) Build(goft *goft.Goft) {
	//goft.Handle("GET", "/userinfo", this.Create)
	goft.Handle("POST", "/accesstoken", AccessToken)
}
