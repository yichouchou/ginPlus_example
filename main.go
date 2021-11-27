package main

import (
	"ginPlus_example/controller"
	_ "ginPlus_example/routers" // Debug mode requires adding [mod] / routes to register annotation routes.debug模式需要添加[mod]/routers 注册注解路由
	"github.com/gin-gonic/gin"
	"github.com/yichouchou/ginPlus/annotation"
)

func main() {
	engine := gin.Default() //todo 考虑兼容 iris的注解路由
	base := annotation.New()
	base.Dev(false)
	base.Register(engine, new(controller.Hello), new(controller.UserRest))
	engine.Run(":8088")
}
