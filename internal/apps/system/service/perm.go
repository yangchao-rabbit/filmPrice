package service

import (
	"filmPrice/internal/apps/system/dao"
	"filmPrice/internal/apps/system/model"
	"filmPrice/internal/models"
	"fmt"
)

func (s *service) PermList(req *model.PermListReq) (*model.PermListResp, error) {
	var (
		total int64
		list  []*dao.SystemPermModel
	)

	db := s.db
	if req.Method != "" {
		db = db.Where("method = ?", req.Method)
	}
	if req.Filter != "" {
		db = db.Where("name like ?", "%"+req.Filter+"%")
	}

	if err := db.Scopes(models.PageOrder(req.Page, req.PageSize)).Find(&list).Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("[Perm] 获取权限列表失败: %v ", err)
	}

	return &model.PermListResp{
		Total: total,
		Rows:  list,
	}, nil
}
