package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/shunjiecloud/encrypt-api/docs"
	v1 "github.com/shunjiecloud/encrypt-api/routers/api/v1"
	"github.com/shunjiecloud/pkg/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/encrypt/v1/")

	url := ginSwagger.URL("https://api.shunjiecloud.com/encrypt/v1/swagger/doc.json") // The url pointing to API definition
	apiv1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	apiv1.GET("publickey", v1.GetPublicKey)

	return r
}
