package setting

import "gopkg.in/ini.v1"

var Conf = new(AppConfig)

type AppConfig struct {
	Port            int `ini:"port"`
	*DatabaseConfig `ini:"database"`
}

type DatabaseConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
