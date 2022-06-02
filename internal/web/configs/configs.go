package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var APPConfig APPConf

// LoadConfig load app's config
// refer from https://www.cnblogs.com/jiujuan/p/13799976.html
func LoadConfig() {
	vyaml := viper.New()
	vyaml.SetConfigName("app")
	vyaml.SetConfigType("yaml")
	vyaml.AddConfigPath("./conf/")

	if err := vyaml.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("load app config read failed, %s)", err.Error()))
	}

	if err := vyaml.Unmarshal(&APPConfig); err != nil {
		panic(fmt.Sprintf("load app config unmarshal failed, %s)", err.Error()))
	}
}

type APPConf struct {
	App  string `yaml:"app"`
	Mode string `yaml:"mode"`
	Port int    `yaml:"port"`
	DB   DB     `yaml:"db"`
	Log  Log    `yaml:"log"`
}

type DB struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbName"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Log struct {
	Level     string `yaml:"level"`
	Formatter string `yaml:"formatter"`
	Dir       string `yaml:"dir"`
}
