package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"kBase-permissions/src/common"
	"kBase-permissions/src/controller"
	"kBase-permissions/src/middleware"
)

func main() {
	goft.Ignite().Config(common.NewDBconfig(), common.NewCasbinconfig()).
		Attach(middleware.NewAuthMiddleware()).
		Mount("", controller.NewAuthController()).Launch()
}
