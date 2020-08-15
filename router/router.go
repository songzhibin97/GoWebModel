/******
** @创建时间 : 2020/8/11 20:49
** @作者 : SongZhiBin
******/
package router

import (
	"GoWebModel/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 注册路由
func SetUp() *gin.Engine {
	if viper.GetString("APP.Mode") == "release" {
		// 发布版本
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(zap.L()), logger.GinRecovery(zap.L(), true))
	// todo 注册自己的业务路由

	return r
}
