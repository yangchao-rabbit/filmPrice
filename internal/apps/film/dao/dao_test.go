package dao

import (
	"filmPrice/config"
	"filmPrice/pkg/logger"
	"testing"
)

func TestTable(t *testing.T) {
	config.LoadConfigYaml("/Users/yangchaoyi/code/go/filmPrice/config.yaml")
	config.Get().Mysql.Init(logger.New(logger.WithDefault()))

	config.GetDB().AutoMigrate(&FilmModel{}, &FilmLinkModel{}, &FilmPriceModel{})
}
