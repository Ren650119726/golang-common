package main

import (
	"fmt"
	"lottery-demo/src/config"
	"lottery-demo/src/core/global"
	"lottery-demo/src/core/gorm"
	"lottery-demo/src/core/log"
	"lottery-demo/src/core/redis"
	"lottery-demo/src/core/server"
	"lottery-demo/src/route"
)

func main() {
	//初始化配置
	global.GVA_CONFIG = config.Init()
	// 初始化zap日志库
	global.GVA_LOG = log.Zap()
	//初始化数据库
	global.GVA_DB = gorm.Gorm()
	// 程序结束前关闭数据库链接
	db, _ := global.GVA_DB.DB()
	defer db.Close()
	//初始化redis
	global.GVA_REDIS = redis.Init()
	//初始化路由
	r := route.NewRoute()
	//初始化http服务
	str, err := server.StartServer(r)
	if err != nil {
		panic(fmt.Sprintf("str:%v,err:%s\n", str, err))
	}
}
