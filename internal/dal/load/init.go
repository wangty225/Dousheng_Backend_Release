package load

import (
	"Dousheng_Backend/utils/config"
	"Dousheng_Backend/utils/zap"
)

// var logger = zap.InitLogger(config.InitConfig("../../../config/logger/dal.yml"))
var logger = zap.InitLogger(config.InitConfig("./config/logger/dal.yml"))
