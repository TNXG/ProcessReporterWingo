package getconf

import (
	"io/ioutil"
	"log"

	"github.com/TNXG/ProcessReporterWingo/tools"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Endpoint string `yaml:"endpoint"`
		Token    string `yaml:"token"`
	} `yaml:"server"`
}

func ReadFile() Config {
	// 读取文件
	workdir, _ := tools.Getwd()
	data, err := ioutil.ReadFile(workdir + "/config.yml")
	if err != nil {
		log.Fatalf("无法读取文件: %v", err)
	}

	// 将YAML反序列化为结构体
	var conf Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("无法解析文件: %v", err)
	}

	// 直接返回
	return conf
}
