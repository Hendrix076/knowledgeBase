package Getter

import (
	"fmt"
	"knowledgeBase/src/common"
	"knowledgeBase/src/models/UserModel"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

type IUserGetter interface {
	GetUserList(userid []int) (user []*UserModel.Userinfo, err error)
	FindByphone(phone string) (user *UserModel.UserModel,err error)
}
type UserGetterImpl struct {
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{}
}

func (this *UserGetterImpl) FindByphone(phone string) (user *UserModel.UserModel,err error) {

	err=common.Gorm.Raw("select * from users where phone =?",phone).First(user).Error
	if err !=nil {
		return nil, err
	}
	//err:=common.Gorm.Where("phone=?", phone).First(&user).Error
	//if err !=nil {
	//	fmt.Println(err)
	//}
	return
}
func (this *UserGetterImpl) GetUserList(userid []int) (users []*UserModel.Userinfo, err error) {
	str := ""
	for _, id := range userid {
		if str != "" {
			str += ","
		}
		str += fmt.Sprintf("%d", id)
	}
	err = common.Gorm.Raw("select id ,name,user_head from users where id in  (" + str + ")").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return
}
