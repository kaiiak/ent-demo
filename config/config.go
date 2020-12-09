package config

import (
	"fmt"
	"net/url"
)

type (
	Config struct {
		Mysql Mysql
	}
	Mysql struct {
		Username  string `yaml:"username" json:"username"  mapstructure:"username"`
		Password  string `yaml:"password" json:"password" mapstructure:"password"`
		Host      string `yaml:"host" json:"host" mapstructure:"host"`
		Port      int    `yaml:"port" json:"port" mapstructure:"port"`
		DBName    string `yaml:"db_name" json:"db_name" mapstructure:"db_name"`
		MaxIdle   int    `yaml:"max_idle" json:"max_idle" mapstructure:"max_idle"`
		MaxConn   int    `yaml:"max_conn" json:"max_conn" mapstructure:"max_conn"`
		LogLevel  string `yaml:"log_level" json:"log_level" mapstructure:"log_level"`
		AutoMerge bool   `yaml:"auto_merge" json:"auto_merge" mapstructure:"auto_merge"`
		Charset   string `yaml:"charset" json:"charset" mapstructure:"charset"`
		Location  string `yaml:"location" json:"location" mapstructure:"location"`
	}
)

func (m Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DBName,
		m.Charset,
		url.QueryEscape(m.Location),
	)
}
