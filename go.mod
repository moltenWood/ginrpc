module github.com/molten/ginrpc

go 1.12

require (
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/validator/v10 v10.4.1
	github.com/goccy/go-json v0.7.0
	github.com/gookit/color v1.5.0
	github.com/json-iterator/go v1.1.9
	github.com/pkg/errors v0.9.1
	go.uber.org/zap v1.21.0
	google.golang.org/genproto v0.0.0-20210207032614-bba0dbe2a9ea // indirect
	google.golang.org/grpc v1.29.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

//replace github.com/molten/ginrpc/public => ./public
