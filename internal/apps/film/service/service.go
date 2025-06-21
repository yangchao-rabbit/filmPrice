package service

import (
	"filmPrice/config"
	"filmPrice/internal/apps"
	"filmPrice/internal/apps/film"
	"filmPrice/internal/apps/film/dao"
	"filmPrice/internal/apps/film/model"
	"gorm.io/gorm"
	"log"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	FilmList(req *model.FilmListReq) (*model.FilmListResp, error)
	FilmDetail(id string) (*dao.FilmModel, error)
	FilmCreate(req *model.FilmCreateReq) error
	FilmUpdate(id string, req *model.FilmUpdateReq) error
	FilmDelete(id string) error

	FilmLinkList(req *model.FilmLinkListReq) (*model.FilmLinkListResp, error)
	FilmLinkDetail(id string) (*dao.FilmLinkModel, error)
	FilmLinkCreate(req *model.FilmLinkCreateReq) error
	FilmLinkUpdate(id string, req *model.FilmLinkUpdateReq) error
	FilmLinkDelete(id string) error

	FilmPriceList(req *model.FilmPriceListReq) (*model.FilmPriceListResp, error)
	FilmPriceDetail(id string) (*dao.FilmPriceModel, error)

	FilmPriceHistoryList(req *model.FilmPriceHistoryListReq) (*model.FilmPriceHistoryListResp, error)
	FilmPriceHistoryDetail(id string) (*dao.FilmPriceHistoryModel, error)
}

type service struct {
	l  *log.Logger
	db *gorm.DB
}

func (*service) i() {}

func (s *service) Name() string {
	return film.AppName
}

func (s *service) Init() error {
	s.l = apps.Log
	s.db = config.GetDB()
	return nil
}

func init() {
	apps.RegistryImpl(&service{})
}
