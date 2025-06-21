package service

import (
	"filmPrice/internal/apps/film/dao"
	"filmPrice/internal/apps/film/model"
	"filmPrice/internal/models"
	"fmt"
)

func (s *service) FilmPriceList(req *model.FilmPriceListReq) (*model.FilmPriceListResp, error) {
	var (
		total int64
		list  []*dao.FilmPriceModel
	)

	db := s.db
	if req.Filter != "" {
		query := "%" + req.Filter + "%"
		db = db.Where("name like ?", query)
	}

	if err := db.Scopes(models.PageOrder(req.Page, req.PageSize)).Find(&list).Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("[FilmPrice] 查询列表失败: %v ", err)
	}

	return &model.FilmPriceListResp{
		Total: total,
		Rows:  list,
	}, nil
}

func (s *service) FilmPriceDetail(id string) (*dao.FilmPriceModel, error) {
	var model dao.FilmPriceModel
	if err := s.db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, fmt.Errorf("[FilmPrice] 查询失败: %v ", err)
	}

	return &model, nil
}
