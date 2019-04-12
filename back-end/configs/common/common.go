package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Common struct {
	EnvModel string `yaml:"envModel"`
	DbAddr   string `yaml:"dbAddr"`
}

var c *Common

func init() {
	f, err := ioutil.ReadFile("/www/configs/catdogs.club/common.yaml")
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		fmt.Println(err)
	}
}

func GetEnvModel() string {
	return c.EnvModel
}

func GetDbAddr() string {
	return c.DbAddr
}
