package load

import (
	"Dousheng_Backend/utils/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type Mysql struct {
	Host     string
	Port     string
	Config   string
	Dbname   string
	Username string
	Password string
}

func InitMysqlConfig() (m *Mysql) {
	logger.Infoln("[mysql]Init mysql from config files")
	globalConfigFile := "./config/yml/mysql.yml"
	//globalConfigFile := "../../../config/yml/mysql.yml"
	logger.Infof("[mysql]globalConfigFile path: %v\n", globalConfigFile)
	_, errGlobal := os.Stat(globalConfigFile)
	var v *viper.Viper
	if os.IsNotExist(errGlobal) {
		logger.Errorln("[mysql]Error: Global config file '%s' not found. ", globalConfigFile)
	} else {
		logger.Infoln("[mysql]加载全局配置文件：'%s'. ", globalConfigFile)
		v = config.InitConfig(globalConfigFile)
	}

	// 获取数据库配置
	mysqlConfig := v.Sub("mysql")
	host := mysqlConfig.GetString("host")
	port := mysqlConfig.GetString("port")
	username := mysqlConfig.GetString("username")
	password := mysqlConfig.GetString("password")
	dbname := mysqlConfig.GetString("dbname")
	config := mysqlConfig.GetString("config")

	mysql := Mysql{
		Host:     host,
		Port:     port,
		Config:   config,
		Dbname:   dbname,
		Username: username,
		Password: password,
	}
	m = &mysql
	return m
}

func (m *Mysql) GetDSN() string {
	var b strings.Builder
	growSize := len(m.Host) + len(m.Port) + len(m.Dbname) + len(m.Username) + len(m.Password) + len(m.Config)
	if len(m.Config) == 0 {
		b.Grow(growSize + 9)
		if _, err := fmt.Fprintf(&b, "%s:%s@tcp(%s:%s)/%s", m.Username, m.Password, m.Host, m.Port, m.Dbname); err != nil {
			return ""
		}
	} else {
		b.Grow(growSize + 10)
		if _, err := fmt.Fprintf(&b, "%s:%s@tcp(%s:%s)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.Dbname, m.Config); err != nil {
			return ""
		}
	}
	dsn := b.String()
	logger.Infof("[mysql]成功获取DSN!%v\n", "dsn")
	return dsn
}
