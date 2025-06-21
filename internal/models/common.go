package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string    `json:"id" gorm:"primaryKey;type:bigint"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func PageOrder(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		db = db.Offset(offset).Limit(pageSize)
		return db.Order("updated_at DESC")
	}
}

func Page(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}

// CustomMap 自定义Map
type CustomMap map[string]interface{}

func (m *CustomMap) Scan(value interface{}) error {
	// 将数据库中的值转换为自定义数据类型
	if value == nil {
		*m = nil
		return nil
	}

	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scan: expected []byte, got %T", value)
	}

	return json.Unmarshal(data, &m)
}

func (m CustomMap) Value() (driver.Value, error) {
	// 将自定义数据类型转换为数据库中的值
	return json.Marshal(m)
}

// Set 设置键值对
func (m *CustomMap) Set(key string, value interface{}) {
	(*m)[key] = value
}

type CustomList []interface{}

func (m *CustomList) Scan(value interface{}) error {
	if value == nil {
		*m = nil
		return nil
	}

	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scan: expected []byte, got %T", value)
	}

	return json.Unmarshal(data, &m)
}

func (m CustomList) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type ListBaseReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Filter   string `json:"filter"`
}
