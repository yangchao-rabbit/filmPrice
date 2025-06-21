package service

import (
	"filmPrice/internal/apps/film/dao"
	"filmPrice/internal/apps/film/model"
	"filmPrice/internal/models"
	"fmt"
	"github.com/jinzhu/copier"
)

func (s *service) FilmList(req *model.FilmListReq) (*model.FilmListResp, error) {
	var (
		total int64
		list  []*dao.FilmModel
	)

	db := s.db
	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
	}
	if req.Brand != "" {
		db = db.Where("brand = ?", req.Brand)
	}
	if req.Format != "" {
		db = db.Where("format = ?", req.Format)
	}
	if req.ISO != "" {
		db = db.Where("iso = ?", req.ISO)
	}
	if req.Filter != "" {
		query := "%" + req.Filter + "%"
		db = db.Where("name like ? or alias = ?", query, query)
	}

	if err := db.Scopes(models.PageOrder(req.Page, req.PageSize)).Find(&list).Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("[Film] 查询列表失败: %v ", err)
	}

	return &model.FilmListResp{
		Total: total,
		Rows:  list,
	}, nil
}

func (s *service) FilmDetail(id string) (*dao.FilmModel, error) {
	var model dao.FilmModel
	if err := s.db.Where("id = ?", id).Preload("Links").Preload("Links.Prices").First(&model).Error; err != nil {
		return nil, fmt.Errorf("[Film] 查询失败: %v ", err)
	}

	return &model, nil
}

func (s *service) FilmCreate(req *model.FilmCreateReq) error {
	var model dao.FilmModel
	if err := s.db.Where("alias = ?", req.Alias).First(&model).Error; err == nil {
		return fmt.Errorf("[Film] 名称已存在: %v ", err)
	}

	if err := copier.Copy(&model, req); err != nil {
		return fmt.Errorf("[Film] 拷贝数据失败: %v ", err)
	}

	if err := s.db.Create(&model).Error; err != nil {
		return fmt.Errorf("[Film] 创建失败: %v ", err)
	}

	return nil
}

func (s *service) FilmUpdate(id string, req *model.FilmUpdateReq) error {
	var model dao.FilmModel
	if err := s.db.Where("id = ?", id).First(&model).Error; err != nil {
		return fmt.Errorf("[Film] 获取失败: %v ", err)
	}

	if err := copier.Copy(&model, req); err != nil {
		return fmt.Errorf("[Film] 拷贝数据失败: %v ", err)
	}

	if err := s.db.Save(&model).Error; err != nil {
		return fmt.Errorf("[Film] 更新失败: %v ", err)
	}

	return nil
}

func (s *service) FilmDelete(id string) error {
	if err := s.db.Where("id = ?", id).Delete(&dao.FilmModel{}).Error; err != nil {
		return fmt.Errorf("[Film] 删除失败: %v ", err)
	}

	return nil
}
