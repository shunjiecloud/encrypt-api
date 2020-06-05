package main

import (
	"log"

	"github.com/micro/go-micro/v2/web"
	"github.com/shunjiecloud/encrypt-api/modules"
	"github.com/shunjiecloud/encrypt-api/routers"
)

// @title 瞬捷云加解密服务
// @version 1.0
// @description 瞬捷云 encrypt-api
// @termsOfService https://www.shunjiecloud.com

// @contact.name zhoushengjie
// @contact.url https://www.shunjiecloud.com
// @contact.email zhou_shengjie@outlook.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host api.shunjiecloud.com
// @BasePath /

func main() {
	//  modules init
	modules.Setup()

	//  Create web
	webSrv := web.NewService(
		web.Name("go.micro.api.encrypt"),
	)

	//  register web handler
	webSrv.Handle("/", routers.InitRouter())

	//  init
	if err := webSrv.Init(); err != nil {
		log.Fatal(err)
	}

	if err := webSrv.Run(); err != nil {
		log.Fatal(err)
	}
}
