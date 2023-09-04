package zap

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"path/filepath"
)

var infoPath = new(string)
var errorPath = new(string)

// InitLogger 初始化zap
func InitLogger(config *viper.Viper) *zap.SugaredLogger {
	*infoPath = config.GetString("info")
	*errorPath = config.GetString("error")
	//规定日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//各级别通用的encoder
	encoder := getEncoder()

	//INFO级别的日志
	infoSyncer := getInfoWriter()
	infoCore := zapcore.NewCore(encoder, infoSyncer, lowPriority)

	//ERROR级别的日志
	errorSyncer := getErrorWriter()
	errorCore := zapcore.NewCore(encoder, errorSyncer, highPriority)

	//合并各种级别的日志
	core := zapcore.NewTee(infoCore, errorCore)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()
	return sugarLogger
}

// InitDefaultLogger 初始化zap
func InitDefaultLogger() *zap.SugaredLogger {
	config := viper.New()
	config.SetConfigFile(filepath.ToSlash("./config/logger/logger.yml"))
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("errno is %+v", err)
	}
	return InitLogger(config)
}

// 自定义日志输出格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 获取INFO的Writer
func getInfoWriter() zapcore.WriteSyncer {
	//lumberJack:日志切割归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   *infoPath,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 获取ERROR的Writer
func getErrorWriter() zapcore.WriteSyncer {
	//lumberJack:日志切割归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   *errorPath,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
