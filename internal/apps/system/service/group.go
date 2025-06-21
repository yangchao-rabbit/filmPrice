package service

import (
	"filmPrice/internal/apps/system/dao"
	"filmPrice/internal/apps/system/model"
	"filmPrice/internal/models"
	"fmt"
	"github.com/jinzhu/copier"
)

func (s *service) GroupList(req *model.GroupListReq) (*model.GroupListResp, error) {
	var (
		total int64
		list  []*dao.SystemGroupModel
	)

	db := s.db
	if req.Filter != "" {
		db = db.Where("name like ?", "%"+req.Filter+"%")
	}

	if err := db.Scopes(models.PageOrder(req.Page, req.PageSize)).Find(&list).Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("[Group] 获取用户组列表失败: %v ", err)
	}

	return &model.GroupListResp{
		Total: total,
		Rows:  list,
	}, nil
}

func (s *service) GroupGet(id string) (*dao.SystemGroupModel, error) {
	var group dao.SystemGroupModel
	if err := s.db.Where("id = ?", id).First(&group).Error; err != nil {
		return nil, fmt.Errorf("[Group] 获取用户组失败: %v ", err)
	}

	return &group, nil
}

func (s *service) GroupCreate(req *model.GroupCreateReq) error {
	var group dao.SystemGroupModel
	if err := s.db.Where("name = ?", req.Name).First(&group).Error; err == nil {
		return fmt.Errorf("[Group] 用户组已存在: %v ", err)
	}

	if err := copier.Copy(&group, req); err != nil {
		return fmt.Errorf("[Group] 复制数据失败: %v ", err)
	}

	if err := s.db.Create(&group).Error; err != nil {
		return fmt.Errorf("[Group] 创建失败: %v ", err)
	}

	return nil
}

func (s *service) GroupUpdate(id string, req *model.GroupUpdateReq) error {
	var group dao.SystemGroupModel
	if err := s.db.Where("id = ?", id).First(&group).Error; err != nil {
		return fmt.Errorf("[Group] 用户组不存在: %v ", err)
	}

	if err := copier.Copy(&group, req); err != nil {
		return fmt.Errorf("[Group] 复制数据失败: %v ", err)
	}

	if err := s.db.Save(&group).Error; err != nil {
		return fmt.Errorf("[Group] 更新失败: %v ", err)
	}

	return nil
}

func (s *service) GroupDelete(id string) error {
	if err := s.db.Delete(&dao.SystemGroupModel{}, id).Error; err != nil {
		return fmt.Errorf("[Group] 删除失败: %v ", err)
	}

	return nil
}
