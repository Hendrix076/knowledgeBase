package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"knowledgeBase/src/common"
	"knowledgeBase/src/configuration"
	"knowledgeBase/src/controllers"
)

func main() {
	goft.Ignite().Config(common.NewDBconfig(),configuration.NewServiceConfig()).Mount("v1",controllers.NewUserController()).Launch()
}
