package service

import (
	"filmPrice/config"
	"filmPrice/internal/apps"
	"filmPrice/internal/apps/task"
	"filmPrice/internal/apps/task/dao"
	"filmPrice/internal/apps/task/model"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	TaskList(req *model.TaskListReq) (*model.TaskListResp, error)
	TaskGet(id string) (*dao.TaskModel, error)
	TaskCreate(req *model.TaskCreateReq) error
	TaskUpdate(id string, req *model.TaskUpdateReq) error
	TaskDelete(id string) error
	TestCron(spec string) ([]string, error)
	TaskRun(id string)

	TaskFuncList() []string
	TaskCurrentCron() interface{}
}

type service struct {
	l  *log.Logger
	db *gorm.DB
}

func (*service) i() {}

func (s *service) Name() string {
	return task.AppName
}

func (s *service) Init() error {
	s.l = apps.Log
	s.db = config.GetDB()
	return nil
}

func init() {
	apps.RegistryImpl(&service{})
}

// TaskLogInit 初始化日志 返回日志ID
func TaskLogInit(req *dao.TaskLogModel) string {
	data := &dao.TaskLogModel{
		ID:     config.Snowflake.Generate().String(),
		TaskID: req.TaskID,
		Name:   req.Name,
	}

	config.GetDB().Model(&dao.TaskLogModel{}).Create(data)

	return data.ID
}

// TaskLogUpdate 更新任务日志的Output 并自动记录时间
func TaskLogUpdate(id string, s string) {
	output := fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), s)
	config.GetDB().Model(&dao.TaskLogModel{}).Where("id = ?", id).Update("output", gorm.Expr("concat(output, ?)", output))
}

func TaskLogStatusUpdate(id string, status string) {
	config.GetDB().Model(&dao.TaskLogModel{}).Where("id = ?", id).Update("status", status)
}
