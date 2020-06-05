package modules

import (
	"github.com/micro/go-micro/v2"
	proto_encrypt "github.com/shunjiecloud-proto/encrypt/proto"
)

type moduleWrapper struct {
	EncryptSrvClient proto_encrypt.EncryptService
}

//ModuleContext 模块上下文
var ModuleContext moduleWrapper

//Setup 初始化Modules
func Setup() {
	//  account-srv client
	m_service := micro.NewService()
	ModuleContext.EncryptSrvClient = proto_encrypt.NewEncryptService("go.micro.srv.encrypt", m_service.Client())
}
