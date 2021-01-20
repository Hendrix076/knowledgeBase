package UserLog

import "time"

type UserLog struct {
	Id        int       `gorm:"column:id"`
	Operator  int       `gorm:"column:operator"`
	Ip        string    `gorm:"column:ip"`
	CreateTime time.Time `gorm:"column:CreateTime"`
}

func NewUserLog(attr...UserLogAttrFunc) *UserLog {
	l:= &UserLog{}
	UserLogAttrFuncs(attr).Apply(l)
	return l
}
func (this *UserLog)Mutate(attr ...UserLogAttrFunc)*UserLog{
	UserLogAttrFuncs(attr).Apply(this)
	return this
}
