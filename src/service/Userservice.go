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
func (this *UserService)Userinfo(userId []int)([]*UserModel.Userinfo,error){

	_,err:=this.UserDao.GetUserinfo(this.DB,userId)
	if err !=nil{
		return nil, nil
	}
	return nil ,nil

}