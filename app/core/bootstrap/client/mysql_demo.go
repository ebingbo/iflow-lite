package client

import (
	"iflow-lite/core/config"
	"iflow-lite/core/mysql"

	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func InitMysqlClient() {
	MysqlDB = mysql.InitMysql(config.Config.MysqlDB)
}
