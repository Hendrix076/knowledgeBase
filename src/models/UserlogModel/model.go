package UserLog

import (
	"fmt"
	"time"
)

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
func (this *UserLog) String() string {
	str := fmt.Sprintf(`{UserId:%d,userOperator:%d,userip:%s,UserLogCreateTime:%s}`,
		this.Id, this.Operator,this.Ip ,this.CreateTime.Format("2006-01-02 15:04:05"))
	return str
}
