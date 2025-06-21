package service

import (
	"filmPrice/internal/apps/film/dao"
	"filmPrice/internal/apps/film/model"
	"filmPrice/internal/models"
	"fmt"
	"github.com/jinzhu/copier"
)

func (s *service) FilmLinkList(req *model.FilmLinkListReq) (*model.FilmLinkListResp, error) {
	var (
		total int64
		list  []*dao.FilmLinkModel
	)

	db := s.db
	if req.Filter != "" {
		query := "%" + req.Filter + "%"
		db = db.Where("name like ?", query)
	}

	if err := db.Scopes(models.PageOrder(req.Page, req.PageSize)).Find(&list).Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("[FilmLink] 查询列表失败: %v ", err)
	}

	return &model.FilmLinkListResp{
		Total: total,
		Rows:  list,
	}, nil
}

func (s *service) FilmLinkDetail(id string) (*dao.FilmLinkModel, error) {
	var model dao.FilmLinkModel
	if err := s.db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, fmt.Errorf("[FilmLink] 查询失败: %v ", err)
	}

	return &model, nil
}

func (s *service) FilmLinkCreate(req *model.FilmLinkCreateReq) error {
	var model dao.FilmLinkModel
	if err := s.db.Where("name = ?", req.Name).First(&model).Error; err == nil {
		return fmt.Errorf("[FilmLink] 名称已存在: %v ", err)
	}

	if err := copier.Copy(&model, req); err != nil {
		return fmt.Errorf("[FilmLink] 拷贝数据失败: %v ", err)
	}

	if err := s.db.Create(&model).Error; err != nil {
		return fmt.Errorf("[FilmLink] 创建失败: %v ", err)
	}

	return nil
}

func (s *service) FilmLinkUpdate(id string, req *model.FilmLinkUpdateReq) error {
	var model dao.FilmLinkModel
	if err := s.db.Where("id = ?", id).First(&model).Error; err != nil {
		return fmt.Errorf("[FilmLink] 获取失败: %v ", err)
	}

	if err := copier.Copy(&model, req); err != nil {
		return fmt.Errorf("[FilmLink] 拷贝数据失败: %v ", err)
	}

	if err := s.db.Save(&model).Error; err != nil {
		return fmt.Errorf("[FilmLink] 更新失败: %v ", err)
	}

	return nil
}

func (s *service) FilmLinkDelete(id string) error {
	if err := s.db.Where("id = ?", id).Delete(&dao.FilmLinkModel{}).Error; err != nil {
		return fmt.Errorf("[FilmLink] 删除失败: %v ", err)
	}

	return nil
}
