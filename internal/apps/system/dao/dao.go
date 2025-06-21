package dao

import (
	"filmPrice/config"
	"filmPrice/internal/models"
	"gorm.io/gorm"
)

type SystemUserModel struct {
	models.BaseModel
	Type     string `json:"type" gorm:"type:varchar(255)"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`

	Groups []*SystemGroupModel `json:"groups" gorm:"many2many:system_user_group;"`

	Desc string `json:"desc" gorm:"type:text"`
}

func (u *SystemUserModel) TableName() string {
	return "system_user"
}

func (u *SystemUserModel) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = config.Snowflake.Generate().String()
	}

	return
}

type SystemGroupModel struct {
	models.BaseModel
	Name string `json:"name" gorm:"type:varchar(255)"`

	Users []*SystemUserModel `json:"users" gorm:"many2many:system_user_group;"`
	Perms []*SystemPermModel `json:"perms" gorm:"many2many:system_group_perm;"`

	Desc string `json:"desc" gorm:"type:text"`
}

func (g *SystemGroupModel) TableName() string {
	return "system_group"
}

func (g *SystemGroupModel) BeforeCreate(tx *gorm.DB) (err error) {
	if g.ID == "" {
		g.ID = config.Snowflake.Generate().String()
	}

	return
}

type SystemPermModel struct {
	models.BaseModel
	Name   string `json:"name" gorm:"type:varchar(255)"`
	Method string `json:"method" gorm:"type:varchar(255)"`
	Path   string `json:"path" gorm:"type:varchar(255)"`
	Desc   string `json:"desc" gorm:"type:text"`
}

func (s *SystemPermModel) TableName() string {
	return "system_perm"
}

func (s *SystemPermModel) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == "" {
		s.ID = config.Snowflake.Generate().String()
	}
	return
}
