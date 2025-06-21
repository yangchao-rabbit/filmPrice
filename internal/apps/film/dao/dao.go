package dao

import (
	"filmPrice/config"
	"filmPrice/internal/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type FilmModel struct {
	models.BaseModel
	// 名称
	Name string `json:"name" gorm:"type:varchar(255)"`
	// 别名
	Alias string `json:"alias" gorm:"type:varchar(255);unique"`
	// 品牌
	Brand string `json:"brand" gorm:"type:varchar(255)"`
	ISO   string `json:"iso" gorm:"type:varchar(255)"`
	// 类型: 彩负 黑白 反转
	Type string `json:"type" gorm:"type:varchar(255)"`
	// 格式: 135 120 45
	Format string `json:"format" gorm:"type:varchar(255)"`
	Image  string `json:"image" gorm:"type:varchar(255)"`
	Desc   string `json:"desc" gorm:"type:text"`

	Links []FilmLinkModel `json:"links" gorm:"foreignKey:FilmID;references:ID;constraint:OnDelete:CASCADE"`
}

func (f *FilmModel) TableName() string {
	return "film"
}

func (f *FilmModel) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == "" {
		f.ID = config.Snowflake.Generate().String()
	}
	return
}

type FilmLinkModel struct {
	models.BaseModel
	FilmID string `json:"film_id" gorm:"type:bigint;index"`
	// 平台
	Platform string `json:"platform" gorm:"type:varchar(255)"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Url      string `json:"url" gorm:"type:varchar(255);unique"`
	IsActive bool   `json:"is_active" gorm:"type:tinyint"`

	Film   FilmModel        `json:"film" gorm:"foreignKey:FilmID;references:ID"`
	Prices []FilmPriceModel `json:"prices" gorm:"foreignKey:LinkID;references:ID;constraint:OnDelete:CASCADE"`
}

func (f *FilmLinkModel) TableName() string {
	return "film_link"
}

func (f *FilmLinkModel) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == "" {
		f.ID = config.Snowflake.Generate().String()
	}
	return
}

type FilmPriceModel struct {
	models.BaseModel
	LinkID string          `json:"link_id" gorm:"type:bigint;index"`
	Price  decimal.Decimal `json:"price" gorm:"type:decimal(10,2)"`
}

func (f *FilmPriceModel) TableName() string {
	return "film_price"
}

func (f *FilmPriceModel) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == "" {
		f.ID = config.Snowflake.Generate().String()
	}
	return
}
