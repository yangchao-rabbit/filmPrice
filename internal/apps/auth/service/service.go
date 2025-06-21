package service

import (
	"filmPrice/config"
	"filmPrice/internal/apps"
	"filmPrice/internal/apps/auth"
	"filmPrice/internal/apps/auth/model"
	"gorm.io/gorm"
	"log"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	// LocalLogin 本地登录
	LocalLogin(req *model.LocalLoginReq) (string, error)
}

type service struct {
	l  *log.Logger
	db *gorm.DB
}

func (*service) i() {}

func (s *service) Name() string {
	return auth.AppName
}

func (s *service) Init() error {
	s.l = apps.Log
	s.db = config.GetDB()
	return nil
}

func init() {
	apps.RegistryImpl(&service{})
}
