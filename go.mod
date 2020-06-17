module github.com/shunjiecloud/encrypt-api

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-gonic/gin v1.6.3
	github.com/micro/go-micro/v2 v2.8.0
	github.com/shunjiecloud-proto/encrypt v0.0.0-20200617093042-b23e2d8c325b
	github.com/shunjiecloud/pkg v0.0.0-20200608213205-7936a725a0c8
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
)
