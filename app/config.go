package app

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var Config = new(Conf)

type Conf struct {
	Env string `yaml:"env"`

	Db     `yaml:"db"`
	Logs   `yaml:"logs"`
	Server `yaml:"server"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SslMode  string `yaml:"sslmode"`
}

type Logs struct {
	Level string `yaml:"level"`
}

type Server struct {
	Port string `yaml:"port"`
}

func LoadConfig(environment string) {
	configFilePath := fmt.Sprintf("app/config/%v.yml", environment)

	yamlFile, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)

	if err != nil {
		log.Fatalln("error while parsing config file", err)
	}
}

func (db *Db) ConnectionString() string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=%v",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
		db.SslMode,
	)
}
