package config

import (
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/sirupsen/logrus"
)

var config *viper.Viper

func Init(env string) {
	config = viper.New()
	config.SetConfigType("yaml")
	viper.SetConfigName(env)
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("fatal error config file: %s \n", err)
	}
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
