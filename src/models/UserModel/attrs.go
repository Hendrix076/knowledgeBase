package UserModel

import "time"

type UserModelAttrFunc  func(u *UserModel)

type UserModelAttrFuncs []UserModelAttrFunc

func (this UserModelAttrFuncs)Apply(u *UserModel){
	for _,f :=range  this {
		f(u)
	}
}
func WithUserID(id int)UserModelAttrFunc{
	return func(u *UserModel) {
		u.Id=id
	}
}
func WithUserName(name string)UserModelAttrFunc{
	return func(u *UserModel) {
		u.Name=name
	}
}
func WithUserPhone(phone string)UserModelAttrFunc{
	return func(u *UserModel) {
		u.Phone=phone
	}
}
func WithUserPassword(Password string)UserModelAttrFunc{
	return func(u *UserModel) {
		u.Password=Password
	}
}
func WithUserStatus(Status int)UserModelAttrFunc{
	return func(u *UserModel) {
		u.Status=Status
	}
}
func WithUserStatusChangeTime(StatusChangeTime time.Time)UserModelAttrFunc{
	return func(u *UserModel) {
		u.StatusChangeTime=StatusChangeTime
	}
}
func WithUserOpenTime(OpenTime time.Time)UserModelAttrFunc{
	return func(u *UserModel) {
		u.OpenTime=OpenTime
	}
}
func WithUserLastLoginTime(LastLoginTime time.Time)UserModelAttrFunc{
	return func(u *UserModel) {
		u.LastLoginTime=LastLoginTime
	}
}