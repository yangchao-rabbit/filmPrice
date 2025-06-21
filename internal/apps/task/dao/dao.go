package dao

import (
	"filmPrice/config"
	"filmPrice/internal/models"
	"gorm.io/gorm"
	"time"
)

type TaskModel struct {
	models.BaseModel
	Name string `json:"name" gorm:"type:varchar(255)"`
	// 任务类型: 手动任务 定时任务
	Type string `json:"type" gorm:"type:varchar(255)"`
	// 定时任务ID
	CronID int `json:"cron_id" gorm:"type:bigint"`
	// 定时表达式
	Cron string `json:"cron" gorm:"type:varchar(255)"`
	// 函数名
	FuncName string `json:"func_name" gorm:"type:varchar(255)"`
	// 参数
	Params models.CustomMap `json:"params" gorm:"type:json"`
	// 状态: 0: 禁用 1: 启用
	IsActive bool   `json:"is_active" gorm:"type:tinyint"`
	Desc     string `json:"desc" gorm:"type:text"`
}

func (m *TaskModel) TableName() string {
	return "task"
}

func (m *TaskModel) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = config.Snowflake.Generate().String()
	}
	return
}

type TaskLogModel struct {
	ID        string `json:"id" gorm:"primary_key;type:bigint"`
	TaskID    string `json:"task_id" gorm:"type:bigint"`
	Name      string `json:"name" gorm:"type:varchar(255)"`
	Status    string `json:"status" gorm:"type:varchar(255)"`
	Output    string `json:"output" gorm:"type:text"`
	CreatedAt time.Time

	Task *TaskModel `json:"task" gorm:"-:migration;foreignKey:TaskID"`
}

func (m *TaskLogModel) TableName() string {
	return "task_log"
}

func (m *TaskLogModel) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		m.ID = config.Snowflake.Generate().String()
	}
	return
}
