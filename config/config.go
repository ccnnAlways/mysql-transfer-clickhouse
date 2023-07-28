package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Conf *Config
)

func init() {
	data, err := os.ReadFile("etc/config.yml")

	if err != nil {
		fmt.Printf("打开配置文件发生异常 = %v \n", err)
		return
	}

	Conf = &Config{}
	if err = yaml.Unmarshal(data, Conf); err != nil {
		fmt.Printf("发序列化发生异常 = %v \n", err)
		return
	}

	fmt.Println("config = ", *Conf)
}

type Config struct {
	Mysql struct {
		Host     string
		Port     int
		User     string
		Passwd   string
		Database string
	} `json:"mysql"`
	Clickhouse struct {
		Host     string
		Port     int
		User     string
		Passwd   string
		Database string
	} `json:"clickhouse"`
}
