package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io"
	"lottery-demo/src/core/common/response"
	"lottery-demo/src/core/global"
	middle_ware "lottery-demo/src/middleware"
	"os"
)

func NewRoute() *gin.Engine {
	//// 禁用控制台颜色，当你将日志写入到文件的时候，你不需要控制台颜色
	gin.DisableConsoleColor()
	// 写入日志文件
	f, _ := os.Create("gin.log")
	gin.SetMode(global.GVA_CONFIG.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.New()
	// 性能分析 - 正式环境不要使用！！！
	if gin.Mode() != gin.ReleaseMode {
		pprof.Register(r)
	}
	//缺失路由
	r.NoRoute(defaultRequest)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "POST", "GET"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}), middle_ware.Recovery(true))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//路由分组
	defaultRouter := r.Group("/")
	{
		defaultRouter.Any("", defaultRequest)
	}
	return r
}

// 默认路由
func defaultRequest(c *gin.Context) {
	response.HttpReturn(c, response.RESPONSE_UNKNOWN, nil, "地址不存在")
	return
}
