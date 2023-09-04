package load

import (
	"Dousheng_Backend/utils/config"
	"github.com/spf13/viper"
	"os"
)

type Redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func InitRedisConfig() (r *Redis) {
	logger.Infoln("[redis]加载配置文件...")

	globalRedisConfigFile := "./config/yml/redis.yml"
	//globalRedisConfigFile := "../../../config/yml/redis.yml"

	_, errRedisGlobal := os.Stat(globalRedisConfigFile)
	var v *viper.Viper
	if os.IsNotExist(errRedisGlobal) {
		logger.Errorf("Global config file '%s' not found. \n", globalRedisConfigFile)
		return
	} else {
		logger.Infof("Global config file '%s' found. \n", globalRedisConfigFile)
		v = config.InitConfig(globalRedisConfigFile)
	}

	// 获取数据库配置
	redisConfig := v.Sub("redis")
	host := redisConfig.GetString("host")
	port := redisConfig.GetString("port")
	password := redisConfig.GetString("password")
	db := redisConfig.GetInt("db")

	redis := Redis{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       db,
	}
	r = &redis
	return
}
