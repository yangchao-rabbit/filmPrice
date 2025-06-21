package service

import (
	"filmPrice/internal/apps/film/dao"
	"filmPrice/internal/apps/film/model"
	"filmPrice/internal/models"
	"fmt"
)

func (s *service) FilmPriceHistoryList(req *model.FilmPriceHistoryListReq) (*model.FilmPriceHistoryListResp, error) {
	var (
		list  []*dao.FilmPriceHistoryModel
		total int64
	)

	db := s.db
	if req.LinkID != "" {
		db = db.Where("link_id = ?", req.LinkID)
	}

	if req.Filter != "" {
		query := "%" + req.Filter + "%"
		db = db.Where("price like ?", query)
	}

	if err := db.Scopes(models.Page(req.Page, req.PageSize)).Order("checked_at desc").Find(&list).Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("[FilmPriceHistory] 获取列表失败: %v ", err)
	}

	return &model.FilmPriceHistoryListResp{
		Rows:  list,
		Total: total,
	}, nil
}

func (s *service) FilmPriceHistoryDetail(id string) (*dao.FilmPriceHistoryModel, error) {
	var model dao.FilmPriceHistoryModel
	if err := s.db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, fmt.Errorf("[FilmPriceHistory] 获取详情失败: %v ", err)
	}

	return &model, nil
}
