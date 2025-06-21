package service

import (
	"filmPrice/config"
	"filmPrice/internal/apps/film/model"
	"filmPrice/pkg/logger"
	"testing"
)

func init() {
	config.LoadConfigYaml("/Users/yangchaoyi/code/go/cmdb/config.yaml")
	config.Get().Mysql.Init(logger.New(logger.WithDefault()))
}

func TestService_List(t *testing.T) {

	svc := &service{
		db: config.GetDB(),
	}

	resp, err := svc.List(&model.CloudListReq{})
	if err != nil {
		return
	}

	t.Log(resp)
}
