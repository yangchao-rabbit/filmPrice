package service

import (
	"filmPrice/config"
	"filmPrice/internal/apps"
	"filmPrice/internal/apps/system"
	"filmPrice/internal/apps/system/dao"
	"filmPrice/internal/apps/system/model"
	"gorm.io/gorm"
	"log"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	UserList(req *model.UserListReq) (*model.UserListResp, error)
	UserGet(id string) (*dao.SystemUserModel, error)
	UserCreate(req *model.UserCreateReq) error
	UserUpdate(id string, req *model.UserUpdateReq) error
	UserDelete(id string) error

	GroupList(req *model.GroupListReq) (*model.GroupListResp, error)
	GroupGet(id string) (*dao.SystemGroupModel, error)
	GroupCreate(req *model.GroupCreateReq) error
	GroupUpdate(id string, req *model.GroupUpdateReq) error
	GroupDelete(id string) error

	PermList(req *model.PermListReq) (*model.PermListResp, error)
}

type service struct {
	l  *log.Logger
	db *gorm.DB
}

func (*service) i() {}

func (s *service) Name() string {
	return system.AppName
}

func (s *service) Init() error {
	s.l = apps.Log
	s.db = config.GetDB()
	return nil
}

func init() {
	apps.RegistryImpl(&service{})
}
