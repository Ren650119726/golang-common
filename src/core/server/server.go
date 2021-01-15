package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lottery-demo/src/core/global"
	"time"
)

var (
	Addr       string
	DeviceToDB time.Duration = 1
)


//启动服务
func StartServer(r *gin.Engine) (string, error) {
	var err error
	Addr = fmt.Sprintf("%s:%d", global.GVA_CONFIG.System.BindIp, global.GVA_CONFIG.System.BindPort)
	if global.GVA_CONFIG.System.UseHttps {
		global.GVA_LOG.Info("HTTPS服务器启动成功")
		if err = r.RunTLS(Addr, global.GVA_CONFIG.System.Cert, global.GVA_CONFIG.System.Key); err != nil {
			return "启动https服务器失败,%s", err
		}
	} else {
		global.GVA_LOG.Info("HTTP服务器启动成功")
		if err = r.Run(Addr); err != nil {
			return "启动http服务器失败,%s", err
		}
	}
	return "", nil
}
