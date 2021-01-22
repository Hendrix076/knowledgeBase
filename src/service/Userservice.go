package service

import (
	"gorm.io/gorm"
	"knowledgeBase/src/dao"
	"knowledgeBase/src/models/UserModel"
)

type UserService struct {
	DB          *gorm.DB         `inject:"-"`
	UserDao    *dao.UserInfoDao    `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}
func (this *UserService)Userinfo(userId []int)(userInfolist *[]UserModel.Userinfo){

	return this.UserDao.GetUserinfo(this.DB,userId)

}