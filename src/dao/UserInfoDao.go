package dao

import (
	"fmt"
	"gorm.io/gorm"

	"knowledgeBase/src/models/UserModel"
)

type UserInfoDao struct{}

func NewUserinfoDao() *UserInfoDao {
	return &UserInfoDao{}
}

func (this *UserInfoDao) GetUserinfo(db *gorm.DB, userid []int) (user []*UserModel.Userinfo, err error) {
fmt.Println(db)
return nil,nil
	//str := ""
	//for _, id := range userid {
	//	if str == "" {
	//		str += ","
	//	}
	//	str += fmt.Sprintf("%d", id)
	//}
	//err = db.Raw("select id,name,user_head from users where id in ("+str+")").Scan(&user).Error
	//
	//return
}
