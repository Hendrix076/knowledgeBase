package service

import (
	"fmt"
	"gorm.io/gorm"
	"knowledgeBase/src/dao"
	"knowledgeBase/src/models/UserModel"
	"log"
)

type UserService struct {
	DB          *gorm.DB         `inject:"-"`
	UserDao    *dao.UserInfoDao
}

func NewUserService() *UserService {
	return &UserService{}
}
func (this *UserService)Userinfo(userId []int)([]*UserModel.Userinfo,error){
	fmt.Println(this.DB)
	//_,err:=this.UserDao.GetUserinfo(this.DB,userId)
	//if err !=nil{
	//	return nil, nil
	//}
	//return nil ,nil
	log.Println("aaaaaaaa")
	return nil, nil
}