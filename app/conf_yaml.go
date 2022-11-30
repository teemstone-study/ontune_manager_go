package app

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DbInfo struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type DemoInfo struct {
	HostCount       int `yaml:"host_count"`
	HostChangeCount int `yaml:"host_change_count"`
	BptCount        int `yaml:"bpt_count"`
	Interval        int `yaml:"interval"`
}

type Config struct {
	Database []DbInfo `yaml:"database"`
	Demo     DemoInfo `yaml:"demo"`
}

func (d DbInfo) Datasource() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Seoul",
		d.Host, d.Port, d.Dbname, d.Username, d.Password)
}

func GetConfig(filename string) Config {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
