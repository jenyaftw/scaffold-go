package config

import (
	"github.com/spf13/viper"
)

type HttpConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver     string `mapstructure:"driver"`
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Name       string `mapstructure:"name"`
	Secure     bool   `mapstructure:"secure"`
	Migrations string `mapstructure:"migrations"`
}

type EmailConfig struct {
	From   string `mapstructure:"from"`
	ApiKey string `mapstructure:"api_key"`
}

type RedisConfig struct {
	Host     string `mapstructure:"addr"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type JwtConfig struct {
	Secret string `mapstructure:"secret"`
}

type Config struct {
	Http  *HttpConfig     `mapstructure:"http"`
	Db    *DatabaseConfig `mapstructure:"db"`
	Redis *RedisConfig    `mapstructure:"redis"`
	Email *EmailConfig    `mapstructure:"email"`
	Jwt   *JwtConfig      `mapstructure:"jwt"`
}

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
