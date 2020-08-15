/******
** @创建时间 : 2020/8/11 20:38
** @作者 : SongZhiBin
******/
package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// rdb:redis全局变量
var rdb *redis.Client

// Init:Redis初始化
func Init() error {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("REDIS.Host"),
			viper.GetInt("REDIS.Port")),
		Password: viper.GetString("REDIS.Password"),
		DB:       viper.GetInt("REDIS.Block"),
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
