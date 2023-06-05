package conf

import (
	cfg "github.com/dollarkillerx/common/pkg/config"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// CONFIG  global configuration
var CONFIG Configuration

// Configuration ...
type Configuration struct {
	ListenAddr       string
	EnablePlayground bool
	JWTConfig        JWTConfiguration
	TimeOut          int
	LoggerConfig     cfg.LoggerConfig

	PostgresConfig cfg.PostgresConfiguration
	Salt           string

	CORSAllowedOrigins []string
	StaticServerConfig StaticServerConfiguration
}

// StaticServerConfiguration 静态资源服务配置
type StaticServerConfiguration struct {
	EntityTagsFileName string
	StaticPath         string
}

var defaultStaticServerConfig = StaticServerConfiguration{
	StaticPath:         "static",
	EntityTagsFileName: ".entity_tags.json",
}

// GetDefault ...
func (s *StaticServerConfiguration) GetDefault() *StaticServerConfiguration {
	return &defaultStaticServerConfig
}

// JWTConfiguration ...
type JWTConfiguration struct {
	SecretKey    string
	OperationKey string
}

// InitConfiguration ...
func InitConfiguration(configName string, configPaths []string) error {
	vp := viper.New()
	vp.SetConfigName(configName)
	vp.AutomaticEnv()
	for _, configPath := range configPaths {
		vp.AddConfigPath(configPath)
	}

	if err := vp.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}

	err := vp.Unmarshal(&CONFIG)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
