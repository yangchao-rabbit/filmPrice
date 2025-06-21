package service

import (
	"filmPrice/config"
	"filmPrice/internal/apps/task/dao"
	"filmPrice/internal/apps/task/model"
	"filmPrice/internal/models"
	"filmPrice/pkg/utils"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/robfig/cron/v3"
	"sort"
	"time"
)

func (s *service) TaskList(req *model.TaskListReq) (*model.TaskListResp, error) {
	var (
		list  []*dao.TaskModel
		total int64
	)

	db := s.db
	if req.Filter != "" {
		query := "%" + req.Filter + "%"
		db = db.Where("name like ? or func_name = ?", query, query)
	}

	if req.Type != "" {
		db = db.Where("type = ?", req.Type)
	}

	if err := db.Scopes(models.PageOrder(req.Page, req.PageSize)).Find(&list).Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return nil, fmt.Errorf("[TaskList] 获取任务列表失败: %v ", err)
	}

	return &model.TaskListResp{
		Total: total,
		Rows:  list,
	}, nil
}

func (s *service) TaskGet(id string) (*dao.TaskModel, error) {
	var task dao.TaskModel
	if err := s.db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, fmt.Errorf("[TaskGet] 获取任务失败: %v ", err)
	}

	return &task, nil
}

func (s *service) TaskCreate(req *model.TaskCreateReq) error {
	var task dao.TaskModel
	if err := s.db.Where("name = ?", req.Name).First(&task).Error; err == nil {
		return fmt.Errorf("[TaskCreate] 名称已存在 ")
	}

	if err := copier.Copy(&task, req); err != nil {
		return fmt.Errorf("[TaskCreate] 复制数据失败: %v ", err)
	}

	if err := s.db.Create(&task).Error; err != nil {
		return fmt.Errorf("[TaskCreate] 创建失败: %v ", err)
	}

	return nil
}

func (s *service) TaskUpdate(id string, req *model.TaskUpdateReq) error {
	var task dao.TaskModel
	if err := s.db.Where("id = ?", id).First(&task).Error; err != nil {
		return fmt.Errorf("[TaskUpdate] 查询失败: %v ", err)
	}

	if err := copier.Copy(&task, req); err != nil {
		return fmt.Errorf("[TaskUpdate] 复制失败: %v ", err)
	}

	// 将CronID置0，保证定时任务自动扫描成功
	cronRemove(&task)
	task.CronID = 0

	if err := s.db.Save(&task).Error; err != nil {
		return fmt.Errorf("[TaskUpdate] 更新失败: %v ", err)
	}

	return nil
}

func (s *service) TaskDelete(id string) error {
	var task dao.TaskModel
	if err := s.db.Where("id = ?", id).First(&task).Error; err != nil {
		return fmt.Errorf("[TaskDelete] 查询失败: %v ", err)
	}

	cronRemove(&task)
	if err := s.db.Delete(&task).Error; err != nil {
		return fmt.Errorf("[TaskDelete] 删除失败: %v ", err)
	}

	return nil
}

// TestCron 测试cron并返回最近10次结果
func (s *service) TestCron(spec string) ([]string, error) {
	parser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	schedule, err := parser.Parse(spec)
	if err != nil {
		return nil, fmt.Errorf("[Task] TestCron 表达式错误: %v ", err)
	}

	now := time.Now()
	result := make([]string, 0, 10)

	for i := 0; i < 10; i++ {
		now = schedule.Next(now)
		result = append(result, now.Format("2006-01-02 15:04:05"))
	}

	return result, nil
}

// TaskFuncList 任务函数列表
func (s *service) TaskFuncList() []string {
	// 获取map中的所有KEY
	result := utils.MapKeys(FuncMap)
	sort.Strings(result)
	return result
}

// TaskCurrentCron 获取当前内存中的定时任务
func (s *service) TaskCurrentCron() interface{} {
	result := map[string]string{}
	for _, v := range config.GetCron().Entries() {
		result["id"] = utils.ToString(v.ID)
		result["next"] = v.Schedule.Next(time.Now()).Format("2006-01-02 15:04:05")
	}

	return result
}
