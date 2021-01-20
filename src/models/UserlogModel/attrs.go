package UserLog

import "time"

type UserLogAttrFunc  func(u *UserLog)

type UserLogAttrFuncs []UserLogAttrFunc

func (this UserLogAttrFuncs)Apply(u *UserLog){
	for _,f :=range  this {
		f(u)
	}
}
func WithUserLogID(id int)UserLogAttrFunc{
	return func(u *UserLog) {
		u.Id=id
	}
}
func WithUserLogOperator (Operator  int)UserLogAttrFunc{
	return func(u *UserLog) {
		u.Operator=Operator
	}
}
func WithUserLogIp(IP string)UserLogAttrFunc{
	return func(u *UserLog) {
		u.Ip=IP
	}
}
func WithUserLogTimestamp(CreateTime time.Time)UserLogAttrFunc{
	return func(u *UserLog) {
		u.CreateTime=CreateTime
	}
}
