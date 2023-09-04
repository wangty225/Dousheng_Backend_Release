package mysql

import (
	"Dousheng_Backend/internal/dal/load"
	"Dousheng_Backend/utils/config"
	"Dousheng_Backend/utils/zap"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DBMysql *gorm.DB
var logger = zap.InitLogger(config.InitConfig("./config/logger/dal.yml"))

// 导入mysql包时会自动执行的函数
// 在 Go 语言中，循环导入包不会导致 init 函数多次调用。
// Go 的编译器会确保每个包的 init 函数只被调用一次，无论是否存在循环导入。
func init() {
	db, err := gorm.Open(mysql.Open(load.InitMysqlConfig().GetDSN()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		panic(fmt.Errorf("failed to open db: %w\n", err))
	}

	// 链路查询优化选项
	if err = db.Use(gormopentracing.New()); err != nil {
		logger.Errorf("register gorm plugin: %v\n", err)
	}

	DBMysql = db
	logger.Infoln("Mysql Connected!")
}
