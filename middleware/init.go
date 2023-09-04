package middleware

import (
	"Dousheng_Backend/utils/config"
	"Dousheng_Backend/utils/zap"
)

var logger = zap.InitLogger(config.InitConfig("./config/logger/middleware.yml"))
