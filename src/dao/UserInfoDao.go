package dao

import (
	"github.com/shenyisyn/goft-gin/goft"
	"gorm.io/gorm"

	"knowledgeBase/src/models/UserModel"
)

type UserInfoDao struct {}

func NewUserinfoDao() *UserInfoDao {
	return &UserInfoDao{}
}

func (this *UserInfoDao)GetUserinfo(db *gorm.DB,userid []int)(user *[]UserModel.Userinfo){
	for _,id:=range userid{
		goft.Error(db.Table("users").Select("id,name,user_head,").Where("id=?",id).First(&user).Error)
	}
	return
}