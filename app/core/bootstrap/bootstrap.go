package bootstrap

import (
	"iflow-lite/core/bootstrap/client"
	"iflow-lite/core/bootstrap/logger"
	"iflow-lite/core/bootstrap/otel"
)

func Init(file string) {
	InitConfig(file)
	otel.InitOTEL()
	client.InitMysqlClient()
	logger.InitLogger()
}
