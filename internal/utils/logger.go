package utils

import (
	cfg "github.com/dollarkillerx/common/pkg/config"
	"github.com/dollarkillerx/common/pkg/logger"
)

// Logger global logger
var Logger *logger.RimeLogger

// InitLogger ...
func InitLogger(loggerConfig cfg.LoggerConfig) {
	Logger = logger.NewLogger().
		Level(loggerConfig.Level.Level()).
		Formatter(logger.DefaultFormatter()).
		Rotation(logger.DefaultRotation(&loggerConfig)).
		SetLogReportCaller(loggerConfig.Level.IsDebugMode())
}
