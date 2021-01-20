package UserModel

import (
	"time"
)

type UserModel struct {
	Id               int       `json:"id"  gorm:"column:id"`
	Name             string    `json:"name" gorm:"column:name" `
	Phone            string    `json:"phone"  gorm:"column:phone"`
	Password         string    `json:"password" gorm:"column:password"`
	Status           int       `json:"status" gorm:"column:status"`
	StatusChangeTime time.Time `json:"status_change_time"  gorm:"column:status_change_time" `
	OpenTime         time.Time `json:"open_time" gorm:"column:open_time"`
	LastLoginTime    time.Time `json:"last_login_time" gorm:"column:last_login_time"`
}

func NewUseModel(attr ...UserModelAttrFunc) *UserModel {

	u := &UserModel{}
	UserModelAttrFuncs(attr).Apply(u)
	return u
}
func (this *UserModel)Mutate(attr ...UserModelAttrFunc)*UserModel{
	UserModelAttrFuncs(attr).Apply(this)
	return this
}
type UserLoginLog struct {
	Id        int       `gorm:"column:id"`
	Operator  int       `gorm:"column:operator"`
	Ip        string    `gorm:"column:ip"`
	timestamp time.Time `gorm:"column:timestamp"`
}
type kbUser struct {
	Id       int       `gorm:"column:id"`
	kbId     int       `gorm:"column:kb_id"`
	UserId   int       `gorm:"column:user_id"`
	JoinTime time.Time `gorm:"column:join_time"`
	CanEdit  string    `gorm:"column:can_edit"`
}
