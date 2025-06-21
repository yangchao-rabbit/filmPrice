package service

import (
	"filmPrice/internal/apps/auth"
	"filmPrice/internal/apps/auth/model"
	"filmPrice/internal/apps/system/dao"
	"filmPrice/pkg/password"
	"fmt"
)

func (s *service) LocalLogin(req *model.LocalLoginReq) (string, error) {
	var user dao.SystemUserModel

	if err := s.db.Where("type = 'local' and name = ?", req.Username).First(&user).Error; err != nil {
		return "", fmt.Errorf("[LocalLogin] 用户不存在: %v ", err)
	}

	if !password.CheckPassword(req.Password, user.Password) {
		return "", fmt.Errorf("[LocalLogin] 密码错误")
	}

	return auth.GenToken(user.Name)
}
