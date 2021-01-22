package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"knowledgeBase/src/common"
	"knowledgeBase/src/models/UserModel"
)


//生成token
func GenerateToken(data map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	for key, value := range data {
		token.Claims.(jwt.MapClaims)[key] = value
	}
	return token.SignedString([]byte("1,2,3"))
}

func GenUserToken(user *UserModel.UserModel)(string,error)  {
	payLoad:=map[string]interface{}{
		"user_id": user.Id,
		"phone":user.Phone,
	}

	if len(user.Password)>0 {
		temp:=sha256.Sum256([]byte(user.Password))
		payLoad["uuid_sha256"]=hex.EncodeToString(temp[:])
	}

	return GenerateToken(payLoad)
}
// 登录校验
func  LoginCheck(user *UserModel.UserModel) (*UserModel.UserModel, error) {

	var u UserModel.UserModel
	// 查询用户
	err := common.Orm.Where("phone = ? ", user.Phone).First(&u).Error
	if err != nil {

		return nil, err
	}

	if u.Password!=user.Password {
		return nil,errors.New("密码不正确")
	}

	return &u, err
}