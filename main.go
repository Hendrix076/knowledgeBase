package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"knowledgeBase/src/controllers"
)

func main() {
	goft.Ignite().Mount("v1",controllers.NewUserController()).Launch()
}
