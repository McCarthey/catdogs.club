package configs

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Common struct {
	EnvModel     string `yaml:"envModel"`
	DbAddr       string `yaml:"dbAddr"`
	SqlLogFile   string `yaml:"sqlLogFile"`
	LogFile      string `yaml:"logFile"`
	IdleConns    int    `yaml:"idleConns"`
	OpenConns    int    `yaml:"openConns"`
	AccessKeyId  string `yaml:"accessKeyID"`
	AccessSecret string `yaml:"accessSecret"`
	SignName     string `yaml:"signName"`
	TemplateCode string `yaml:"templateCode"`
	FromEmail    string `yaml:"fromEmail"`
	EmailPasswd  string `yaml:"emailPasswd"`
	ActivateUrl  string `yaml:"activateUrl"`
	AesKey       string `yaml:"aesKey"`
	PwSalt       string `yaml:"pwSalt"`
}

var C *Common

func init() {
	f, err := ioutil.ReadFile("/www/configs/catdogs.club/common.yaml")
	err = yaml.Unmarshal(f, &C)
	if err != nil {
		fmt.Println(err)
	}
}
