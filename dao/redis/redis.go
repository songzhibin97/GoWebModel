/******
** @创建时间 : 2020/8/11 20:38
** @作者 : SongZhiBin
******/
package redis

import (
	"GoWebModel/settings"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// rdb:redis全局变量
var rdb *redis.Client

// Init:Redis初始化
func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			settings.GetString("REDIS.Host"),
			settings.GetInt("REDIS.Port")),
		Password: settings.GetString("REDIS.Password"),
		DB:       settings.GetInt("REDIS.Block"),
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		zap.L().Error(fmt.Sprintf("Redis No Connect Error:%s", err))
		return err
	}
	return nil
}

func Close() {
	_ = rdb.Close()
}
