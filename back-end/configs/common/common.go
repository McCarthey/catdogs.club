package configs

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	EnvModel     string
	DbAddr       string
	SqlLogFile   string
	IdleConns    int
	OpenConns    int
	AccessKeyId  string
	AccessSecret string
	SignName     string
	TemplateCode string
)

type Common struct {
	EnvModel     string `yaml:"envModel"`
	DbAddr       string `yaml:"dbAddr"`
	SqlLogFile   string `yaml:"sqlLogFile"`
	IdleConns    int    `yaml:"idleConns"`
	OpenConns    int    `yaml:"openConns"`
	AccessKeyId  string `yaml:"accessKeyID"`
	AccessSecret string `yaml:"accessSecret"`
	SignName     string `yaml:"signName"`
	TemplateCode string `yaml:"templateCode"`
}

var c *Common

func init() {
	f, err := ioutil.ReadFile("/www/configs/catdogs.club/common.yaml")
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		fmt.Println(err)
	}

	initFields()
}

func initFields() {
	EnvModel = c.EnvModel
	DbAddr = c.DbAddr
	SqlLogFile = c.SqlLogFile
	IdleConns = c.IdleConns
	OpenConns = c.OpenConns
	AccessKeyId = c.AccessKeyId
	AccessSecret = c.AccessSecret
	SignName = c.SignName
	TemplateCode = c.TemplateCode
}