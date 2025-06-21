package service

import (
	"filmPrice/internal/apps/system/dao"
	"filmPrice/internal/apps/system/model"
	"filmPrice/internal/models"
	"filmPrice/pkg/password"
	"fmt"
	"github.com/jinzhu/copier"
)

func (s *service) UserList(req *model.UserListReq) (*model.UserListResp, error) {
	var (
		total int64
		list  []*dao.SystemUserModel
	)

	db := s.db
	if req.Filter != "" {
		db = db.Where("name like ?", "%"+req.Filter+"%")
	}

	if err := db.Scopes(models.PageOrder(req.Page, req.PageSize)).Preload("Groups").Find(&list).Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("[User] 获取用户列表失败: %v ", err)
	}

	return &model.UserListResp{
		Total: total,
		Rows:  list,
	}, nil
}

func (s *service) UserGet(id string) (*dao.SystemUserModel, error) {
	var user dao.SystemUserModel
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, fmt.Errorf("[User] 获取用户失败: %v ", err)
	}
	return &user, nil
}

func (s *service) UserCreate(req *model.UserCreateReq) error {
	var model dao.SystemUserModel
	if err := s.db.Where("name = ?", req.Name).First(&model).Error; err == nil {
		return fmt.Errorf("[User] 用户名已存在: %v ", err)
	}

	if err := copier.Copy(&model, req); err != nil {
		return fmt.Errorf("[User] 复制数据失败: %v ", err)
	}

	if req.Type == "Local" {
		encrypt, err := password.GenPassword(model.Password)
		if err != nil {
			return fmt.Errorf("[User] 加密失败: %v ", err)
		}

		model.Password = encrypt
	}

	if err := s.db.Create(&model).Error; err != nil {
		return fmt.Errorf("[User] 创建失败: %v ", err)
	}

	return nil
}

func (s *service) UserUpdate(id string, req *model.UserUpdateReq) error {
	var model dao.SystemUserModel
	if err := s.db.Where("id = ?", id).First(&model).Error; err != nil {
		return fmt.Errorf("[User] 查询失败: %v ", err)
	}

	if req.Type == "Local" && req.Password != model.Password {
		encrypt, err := password.GenPassword(model.Password)
		if err != nil {
			return fmt.Errorf("[User] 加密失败: %v ", err)
		}

		req.Password = encrypt
	}

	if err := copier.Copy(&model, req); err != nil {
		return fmt.Errorf("[User] 复制数据失败: %v ", err)
	}

	if err := s.db.Save(&model).Error; err != nil {
		return fmt.Errorf("[User] 更新失败: %v ", err)
	}

	return nil
}

func (s *service) UserDelete(id string) error {
	if err := s.db.Where("id = ?", id).Delete(&dao.SystemUserModel{}).Error; err != nil {
		return fmt.Errorf("[User] 删除失败: %v ", err)
	}

	return nil
}
