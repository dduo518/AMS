package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"AMS/src/utils"
)

//定义conf类型
//类型里的属性，全是配置文件里的属性
type conf struct {
	GO_ENV string `yaml:"GO_ENV"`
	PORT   string `yaml:"PORT"`
	PWD    string `yaml:"PWD"`
	Redis  redis  `yaml:"REDIS"`
	Mysql  mysql  `yaml:"MYSQL"`
	TempDir string
	Letter letter `yaml:"LETTER"`
	GprcConf gprcConf `yaml:"GRPC"`
	TokenSecret string `yaml:"TOKENSECRET"`
	Cos cos `yaml:"COS"`
}

type cos struct {
	Host     string `yaml:"HOST"`
	Name     string `yaml:"NAME"`
	SecretID     string `yaml:"SECRETID"`
	SecretKey    string `yaml:"SECRETKEY"`
	AppId    string `yaml:"APPID"`
	Bucket    string `yaml:"BUCKET"`
	Region    string `yaml:"REGION"`
}

type redis struct {
	Host     string `yaml:"HOST"`
	Password string `yaml:"PASSWORD"`
	DB       string `yaml:"DB"`
}

type mysql struct {
	Host     string `yaml:"HOST"`
	Port     string `yaml:"PORT"`
	Username string `yaml:"USERNAME"`
	Password string `yaml:"PASSWORD"`
	Database string `yaml:"DATABASE"`
}

type letter struct {
	AppId string `yaml:"APPID"`
	AppKey string `yaml:"APPKEY"`
	TplId string `yaml:"TPLID"`
	Sign string `yaml:"SIGN"`
	Timeout string `yaml:"TIMEOUT"`
	Api map[string]string `yaml:"API"`
}
type gprcConf struct {
	Port string `yaml:"PORT"`
	Network string `yaml:"NETWOK"`
}
var Conf conf

func init() {
	fileDir := util.GetCurrentPath()
	confName := flag.String("config", "conf.dev.yaml", "string")
	flag.Parse()
	filePath := filepath.Join(fileDir,"../config", *confName)
	Conf.TempDir =filepath.Join(fileDir,"../")
	Conf.GetConf(filePath)
	fmt.Println(Conf)
}

//读取Yaml配置文件,
func (c *conf) GetConf(dir string) *conf {
	yamlFile, err := ioutil.ReadFile(dir)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
