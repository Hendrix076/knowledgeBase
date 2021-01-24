package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"io/ioutil"
	"knowledgeBase/src/models/UserModel"
	"log"
	"time"
)

type TokenCheck struct {
	Phone    string `json:"phone"`
	jwt.StandardClaims
}
func NewTokenCheck() *TokenCheck {
	return &TokenCheck{}
}

func (this *TokenCheck) OnRequest(ctx *gin.Context) error {
	if ctx.Query("token") == "" {
		goft.Throw("token requred", 503, ctx)
	}
	return nil
}
//生成token
func GenerateToken(Phone string) string {
	claims := &TokenCheck{
		Phone: Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(25 * time.Second).Unix(),  //默认25秒过期时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	prikeyBytes, err := ioutil.ReadFile("./pem/private.pem")
	if err != nil {
		log.Fatal("私钥读取失败")
	}
	prikey, err := jwt.ParseRSAPrivateKeyFromPEM(prikeyBytes)
	if err != nil {
		log.Fatal("私钥文件不正确")
	}
	token_string,err:=token.SignedString(prikey)
	if err!=nil{
		log.Println(err)
		return ""
	}
	return token_string
}

//验证token
func  ParseUserToken(user *UserModel.UserModel) (*TokenCheck, error) {
	claims := &TokenCheck{}
	token:=GenerateToken(claims.Phone)
	publickeyBytes, err := ioutil.ReadFile("./pem/public.pem")
	if err != nil {
		log.Fatal("公钥读取失败")
	}
	publickey, err := jwt.ParseRSAPublicKeyFromPEM(publickeyBytes)
	if err != nil {
		log.Fatal("公钥文件不正确")
	}
	_,err= jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return publickey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("错误的 token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New(" 过期的 token ")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token 未激活")
			} else {
				return nil, errors.New("token无效")
			}
		}
	}
	return claims,nil
}