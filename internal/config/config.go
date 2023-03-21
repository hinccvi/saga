package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string `mapstructure:"name"`
		Cert string `mapstructure:"cert"`
		Key  string `mapstructure:"key"`
	} `mapstructure:"app"`

	Context struct {
		Timeout int `mapstructure:"timeout"`
	} `mapstructure:"context"`

	Jwt struct {
		AccessSigningKey  string `mapstructure:"access_signing_key"`
		AccessExpiration  int    `mapstructure:"access_expiration"`
		RefreshSigningKey string `mapstructure:"refresh_signing_key"`
		RefreshExpiration int    `mapstructure:"refresh_expiration"`
	} `mapstructure:"jwt"`

	Postgres struct {
		Dsn string `mapstructure:"dsn"`
	} `mapstructure:"postgres"`

	Kafka struct {
		Host string `mapstructure:"host"`
	} `mapstructure:"kafka"`

	Mongo struct {
		Dsn string `mapstructure:"dsn"`
	} `mapstructure:"mongo"`
}

func Load(service, env string) (Config, error) {
	file := env

	viper.SetConfigName(file)
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("../../config/")
	viper.AddConfigPath("../../../config/")

	conf := new(Config)

	if err := viper.ReadInConfig(); err != nil {
		return *conf, err
	}

	if err := viper.Unmarshal(&conf); err != nil {
		return *conf, err
	}

	return *conf, nil
}
