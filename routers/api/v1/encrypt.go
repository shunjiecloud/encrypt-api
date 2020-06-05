package v1

import (
	"context"

	"github.com/gin-gonic/gin"
	proto_encrypt "github.com/shunjiecloud-proto/encrypt/proto"
	"github.com/shunjiecloud/encrypt-api/modules"
	"github.com/shunjiecloud/encrypt-api/schemas"
	"github.com/shunjiecloud/pkg/app"
)

//
// @Summary 获取公钥
// @tags encrypt
// @Description 获取公钥。
// @Produce  json
// @Success 200 {object} schemas.GetPublicKeyResponse string	"ok"
// @Router /encrypt/v1/publickey [get]
func GetPublicKey(c *gin.Context) {
	//  ctx
	appCtx := app.AppContext{
		GinCtx: c,
	}

	getPublicKeyResp, err := modules.ModuleContext.EncryptSrvClient.GetPublicKey(context.Background(), &proto_encrypt.GetPublicKeyRequest{})
	if err != nil {
		appCtx.WriteError(err)
		return
	}
	var ret schemas.GetPublicKeyResponse
	ret.PublicKey = getPublicKeyResp.PublicKey
	appCtx.WriteJson(200, &ret)
}
