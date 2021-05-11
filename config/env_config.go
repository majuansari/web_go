package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

//EnvConfig struct
type EnvConfig struct {
	// Auth AuthConfig
	DB     DBConfig
	Cache  CacheConfig
	Tracer TracerConfig
	// HTTP HTTPConfig
}

type TracerConfig struct {
	Provider    string  `mapstructure:"TRACER_PROVIDER"`
	ReporterUri string  `mapstructure:"TRACER_REPORTER_URI"`
	ServiceName string  `mapstructure:"TRACER_SERVICE_NAME"`
	Probability float64 `mapstructure:"TRACER_PROBABILITY"`
}

type CacheConfig struct {
	Driver string `mapstructure:"CACHE_DRIVER"`
	Host   string `mapstructure:"CACHE_HOST"`
	Port   string `mapstructure:"CACHE_PORT"`
}
type DBConfig struct {
	Host           string        `mapstructure:"DB_HOST"`
	Port           int           `mapstructure:"DB_PORT"`
	User           string        `mapstructure:"DB_USER"`
	Driver         string        `mapstructure:"DB_DRIVER"`
	Password       string        `mapstructure:"DB_PASSWORD"`
	Database       string        `mapstructure:"DB_DATABASE"`
	MaxIdleConns   int           `mapstructure:"DB_MAXIDLECONNS"`
	MaxOpenConns   int           `mapstructure:"DB_MAXOPENCONNS"`
	ConMaxLifeTime time.Duration `mapstructure:"DB_CONMAXLIFETIME"`
}

//NewEnvConfig constructor
func NewEnvConfig() *EnvConfig {
	var (
		dbConfig     DBConfig
		cacheConfig  CacheConfig
		tracerConfig TracerConfig
	)

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.MergeInConfig()
	if err != nil {
		panic(err)
	}
	bindEnv(&dbConfig)
	bindEnv(&cacheConfig)
	bindEnv(&tracerConfig)
	return &EnvConfig{
		DB:     dbConfig,
		Cache:  cacheConfig,
		Tracer: tracerConfig,
	}
}

// Dialect is define name database
func (cfg DBConfig) Dialect() string {
	return "mysql"
}

// ConnectionInfo get connection info
func (cfg DBConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

func bindEnv(cfg interface{}) {
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
}
