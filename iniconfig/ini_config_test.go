package iniconfig

import (
	"testing"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Password string  `ini:"password"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float64 `ini:"timeout"`
}

func TestIniConfig(t *testing.T) {
	fmt.Println("hello")
	bytes, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error(err)
	}

	conf := new(Config)
	err = UnMarshal(bytes, conf)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Unmarshal success, conf:%#v", conf)
}
