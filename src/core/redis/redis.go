package redis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"lottery-demo/src/core/global"
)

func Init() (client *redis.Client) {
	redisCfg := global.GVA_CONFIG.Redis
	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.GVA_LOG.Info("redis启动成功")
	}
	return
}
