package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db" default:0`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
