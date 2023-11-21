package config

import (
	"strings"

	"github.com/phinc275/gfas/pkg/eventstroredb"
	"github.com/phinc275/gfas/pkg/logger"
	"github.com/spf13/viper"
)

type Config struct {
	AppName          string
	EventStoreConfig eventstroredb.EventStoreConfig `mapstructure:"eventStoreConfig"`
	Http             Http                           `mapstructure:"http"`
	Logger           *logger.Config                 `mapstructure:"logger"`
}

type Http struct {
	Addr            string   `mapstructure:"addr" validate:"required"`
	Development     bool     `mapstructure:"development"`
	BasePath        string   `mapstructure:"basePath" validate:"required"`
	AchievementPath string   `mapstructure:"achievementPath" validate:"required"`
	Origins         []string `mapstructure:"origins"`
}

func DefaultConfig() *Config {
	return &Config{
		AppName: "Achievement System",
		EventStoreConfig: eventstroredb.EventStoreConfig{
			ConnectionString: "esdb://eventstoredb:2113?tls=false",
		},
		Http: Http{
			Addr:            ":8088",
			Development:     true,
			BasePath:        "/api/v1",
			AchievementPath: "/achievements",
			Origins:         make([]string, 0),
		},
		Logger: &logger.Config{
			LogLevel: "info",
			DevMode:  true,
			Encoder:  "console",
		},
	}
}

func InitConfig() (*Config, error) {
	cfg := DefaultConfig()

	viper.SetEnvPrefix("gfas")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.MustBindEnv("eventStoreConfig.connectionString")
	viper.MustBindEnv("http.addr")
	viper.MustBindEnv("http.development")
	viper.MustBindEnv("http.basePath")
	viper.MustBindEnv("http.achievementPath")
	viper.MustBindEnv("http.origins")
	viper.MustBindEnv("logger.level")
	viper.MustBindEnv("logger.devMode")
	viper.MustBindEnv("logger.encoder")

	err := viper.Unmarshal(cfg)
	return cfg, err
}
