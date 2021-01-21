package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (this *UserController) Name() string {
	return "UserController"
}

func (this *UserController) Create(c *gin.Context) string {
	users:=c.Query("users")
	return "success"
}


func (this *UserController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/userinfo", this.Create)
	goft.Handle("POST","/accesstoken",this.accesstoken)
}