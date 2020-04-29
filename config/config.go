package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Type         string `yaml:"type"`
		Name         string `yaml:"name"`
		Host 		 string `yaml:"host"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		MaxIdleConns int `yaml:"maxIdleConns"`
	} `yaml:"database"`
	Cookie struct{
		Path string `yaml:"path"`
		Domain string `yaml:"domain"`
		MaxAge int`yaml:"maxAge"`
		Secure bool`yaml:"secure"`
		HttpOnly bool`yaml:"httpOnly"`
	} `yaml:"cookie"`
	WeChat struct{
		AppId string `yaml:"appId"`
		AppSecret string `yaml:"appSecret"`
		RedirectURL string `yaml:"redirectURL"`
	} `yaml:"weChat"`
}

var c *config

func InitConfig(path string) error {
	b,err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b,&c)
}

func GetConfig() *config {
	return c
}