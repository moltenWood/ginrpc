package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/molten/ginrpc"
)

type ReqTest struct {
	Access_token string `json:"access_token"`
	UserName     string `json:"user_name" binding:"required"` // With verification mode
	Password     string `json:"password"`
}

// Hello ...
type Hello struct {
}

// Hello Annotated route (bese on beego way)
// @Router /block [post,get]
func (s *Hello) Hello(c *CtrlContext, req *ReqTest) {
	fmt.Println(c.Db)
	c.Ctx.JSON(http.StatusOK, "ok")
}

// Hello2 Route without annotation (the parameter is 2 default post)
func (s *Hello) Hello2(c *gin.Context, req ReqTest) {
	fmt.Println(req)
	c.JSON(http.StatusOK, "ok")
}

// [grpc-go](https://github.com/grpc/grpc-go)
// with request,return parameter and error
// Fun6 Route without annotation (the parameter is 2 default post)
func Fun6(c *gin.Context, req ReqTest) (*ReqTest, error) {
	fmt.Println(req)
	//c.JSON(http.StatusOK, req)
	return &req, nil
}

type AppModuleInitFunction func(domain string, e ApiEngineInterface)

type ApiEngineInterface interface {
	GetBusinessDb(domain string) string
	GetAllBusinessDb() map[string]string
	AddBusinessDb(domainid, login_id, domain, datasource string, sync_db bool) bool
	AppModuleInit(domain string)
}

type CtrlContext struct {
	App              ApiEngineInterface
	Ctx              *gin.Context
	Db               string
	PlatformDbEngine string
}

func myContext(c *gin.Context) interface{} {
	println("11111111")
	var myC CtrlContext
	myC.Db = "myDb"
	myC.Ctx = c
	return &myC
}

func TestRun(t *testing.T) {

	// swagger
	//myswagger.SetHost("https://localhost:8080")
	//myswagger.SetBasePath("gmsec")
	//myswagger.SetSchemes(true, false)
	// -----end --
	base := ginrpc.New(ginrpc.WithCtx(myContext))
	router := gin.Default()           // or router :=  gin.Default().Group("/xxjwxc")
	base.Register(router, new(Hello)) // object register like(go-micro)
	//router.POST("/test6", base.HandlerFunc(Fun6))                            // function register
	//base.RegisterHandlerFunc(router, []string{"post", "get"}, "/test", Fun6)
	router.Run(":8080")
}
