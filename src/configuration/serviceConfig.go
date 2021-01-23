package configuration

import (
	"knowledgeBase/src/dao"
	"knowledgeBase/src/service"
)

type ServiceConfig struct{}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (this *ServiceConfig) OrderDao()*dao.UserInfoDao  {
	return dao.NewUserinfoDao()
}

func (this *ServiceConfig) OrderService() *service.UserService {
	return service.NewUserService()
}

