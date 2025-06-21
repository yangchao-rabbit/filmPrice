package config

import (
	"github.com/bwmarrin/snowflake"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func init() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Println(err)
	}

	Snowflake = node

	cr = cron.New()
	cr.Start()
}

// LoadConfigYaml 加载配置文件
//
// 不传参：默认从 当前目录、etc目录下读取 config.yaml文件
//
// 传参：只读取第一个参数作为文件路径
func LoadConfigYaml(filePath ...string) {

	vp := viper.New()
	// 自动注入环境变量
	vp.AutomaticEnv()
	// 设置环境变量前缀
	vp.SetEnvPrefix("vp")
	replacer := strings.NewReplacer(".", "_")
	vp.SetEnvKeyReplacer(replacer)

	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	if len(filePath) >= 1 {
		vp.SetConfigFile(filePath[0])
	} else {
		vp.AddConfigPath(".")
		vp.AddConfigPath("etc")
		vp.AddConfigPath("../etc")
	}

	err := vp.ReadInConfig()
	if err != nil {
		log.Panicf("读取配置文件异常：%v", err)
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		log.Panicf("解析配置文件异常：%v", err)
	}
}
