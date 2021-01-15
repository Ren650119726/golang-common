package config

import (
	"github.com/jinzhu/configor"
)

type Config struct {
	APPName     string `default:"app name"`
	ReleaseMode string `mapstructure:"release-mode" json:"releaseMode" yaml:"release-mode"`
	System      System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql       Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap         Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis       Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
}

type System struct {
	BindIp   string `mapstructure:"bind-ip" json:"bindIp" yaml:"bind-ip"`
	BindPort uint32 `mapstructure:"bind-port" json:"bindPort" yaml:"bind-port"`
	UseHttps bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`
	Cert     string `mapstructure:"cert" json:"cert" yaml:"cert"`
	Key      string `mapstructure:"key" json:"key" yaml:"key"`
}

func Init() Config {
	var conf = Config{}
	err := configor.Load(&conf, "config.yml")
	if err != nil {
		panic(err)
	}
	return conf
}
