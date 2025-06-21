package dao

import (
	"cmdb/config"
	"cmdb/pkg/logger"
	"testing"
)

func TestTable(t *testing.T) {
	config.LoadConfigYaml("/Users/yangchaoyi/code/go/cmdb/config.yaml")
	config.Get().Mysql.Init(logger.New(logger.WithDefaults()))

	config.GetDB().AutoMigrate(&SystemUserModel{}, &SystemGroupModel{}, &SystemPermModel{})
}
