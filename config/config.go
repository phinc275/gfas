package config

import (
	"strings"

	"github.com/phinc275/gfas/pkg/eventstroredb"
	"github.com/phinc275/gfas/pkg/logger"
	"github.com/phinc275/gfas/pkg/mq/kafka"
	"github.com/spf13/viper"
)

type Config struct {
	AppName          string
	EventStoreConfig eventstroredb.EventStoreConfig `mapstructure:"eventStoreConfig"`
	Http             Http                           `mapstructure:"http"`
	Logger           *logger.Config                 `mapstructure:"logger"`
	Kafka            *kafka.Config                  `mapstructure:"kafka"`
	EventHandler     *EventHandler                  `mapstructure:"eventHandler"`
}

type Http struct {
	Addr            string   `mapstructure:"addr" validate:"required"`
	Development     bool     `mapstructure:"development"`
	BasePath        string   `mapstructure:"basePath" validate:"required"`
	AchievementPath string   `mapstructure:"achievementPath" validate:"required"`
	Origins         []string `mapstructure:"origins"`
	JWKS            string   `mapstructure:"jwks"`
}

type EventHandler struct {
	Topics       []string `mapstructure:"topics"`
	NumOfWorkers int      `mapstructure:"numOfWorkers"`
}

func DefaultConfig() *Config {
	return &Config{
		AppName: "Achievement System",
		EventStoreConfig: eventstroredb.EventStoreConfig{
			ConnectionString: "esdb://localhost:2113?tls=false",
		},
		Http: Http{
			Addr:            ":8088",
			Development:     true,
			BasePath:        "/api/v1",
			AchievementPath: "/achievements",
			Origins:         make([]string, 0),
			JWKS:            "http://localhost:8080/auth/jwks",
		},
		Logger: &logger.Config{
			LogLevel: "info",
			DevMode:  true,
			Encoder:  "console",
		},
		Kafka: &kafka.Config{
			BootstrapServers:  "localhost:9092",
			GroupID:           "achievementsystem",
			AutoOffsetReset:   "earliest",
			SchemaRegistryURL: "http://localhost:8081",
		},
		EventHandler: &EventHandler{
			Topics:       []string{"events"},
			NumOfWorkers: 100,
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
	viper.MustBindEnv("http.jwks")
	viper.MustBindEnv("logger.level")
	viper.MustBindEnv("logger.devMode")
	viper.MustBindEnv("logger.encoder")
	viper.MustBindEnv("kafka.bootstrapServers")
	viper.MustBindEnv("kafka.groupID")
	viper.MustBindEnv("kafka.autoOffsetReset")
	viper.MustBindEnv("kafka.schemaRegistryURL")
	viper.MustBindEnv("eventHandler.topics")
	viper.MustBindEnv("eventHandler.numOfWorkers")

	err := viper.Unmarshal(cfg)
	return cfg, err
}
