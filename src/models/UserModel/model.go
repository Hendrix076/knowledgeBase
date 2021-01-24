package UserModel

import (
	"fmt"
	"time"
)

type UserModel struct {
	Id               int       `json:"user_id"  gorm:"column:id"`
	Name             string    `json:"user_name" gorm:"column:name" `
	Phone            string    `json:"phone" gorm:"column:phone"`
	Password         string    `json:"password" gorm:"column:password"`
	UserHead         string    `json:"user_head" gorm:"column:name" `
	Status           int       `json:"status" gorm:"column:status"`
	StatusChangeTime time.Time `json:"status_change_time"  gorm:"column:status_change_time" `
	OpenTime         time.Time `json:"open_time" gorm:"column:open_time"`
	LastLoginTime    time.Time `json:"last_login_time" gorm:"column:last_login_time"`
}
type ResUserInfo struct {
	Users []int `json:"users"`
}

type Userinfo struct {
	UserId   int    `json:"user_id"  gorm:"column:id"   `
	Username string `json:"user_name" gorm:"column:name" `
	UserHead string `json:"user_head" gorm:"column:user_head" `
}

func NewUseModel(attr ...UserModelAttrFunc) *UserModel {

	u := &UserModel{}
	UserModelAttrFuncs(attr).Apply(u)
	return u
}
func (this *UserModel) String() string {
	str := fmt.Sprintf(`{UserId:%d,UserName:%s,UserPhone:%s,UserPassword:%s,userStatus:%d,StatusChangeTime:%s,OpenTime:%s,LastLoginTime:%s}`,
		this.Id, this.Name, this.Phone, this.Password, this.Status, this.StatusChangeTime.Format("2006-01-02 15:04:05"), this.OpenTime.Format("2006-01-02 15:04:05"), this.LastLoginTime.Format("2006-01-02 15:04:05"))
	return str
}
func (this *UserModel) Mutate(attr ...UserModelAttrFunc) *UserModel {
	UserModelAttrFuncs(attr).Apply(this)
	return this
}
