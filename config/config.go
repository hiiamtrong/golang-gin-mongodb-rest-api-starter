package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppEnv string

const (
	AppEnvDebug   AppEnv = "debug"
	AppEnvRelease AppEnv = "release"
	AppEnvTest    AppEnv = "test"
)

func (a AppEnv) String() string {
	return string(a)
}

type Config struct {
	App struct {
		Env AppEnv
	}
	Server struct {
		Port         string
		Host         string
		URL          string
		ReadTimeout  string
		WriteTimeout string
	}
	Swagger struct {
		Host     string
		BasePath string
		Schemes  string
		Username string
		Password string
	}
	Log struct {
		Level string
	}
	Mongodb struct {
		Host     string
		Port     string
		Database string
		Username string
		Password string
	}

	JWT struct {
		Secret     string
		Expiration string
	}
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config.NewConfig: load file failed: %w", err)
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("config.NewConfig: unmarshal to struct failed: %w", err)
	}

	if cfg.App.Env != AppEnvRelease {
		printJSONPretty(&cfg)
	}

	return &cfg, nil
}

func printJSONPretty(cfg *Config) {
	var prettyJSON []byte
	prettyJSON, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Current config:\n%s\n", string(prettyJSON))
}
