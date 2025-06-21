package service

import (
	"filmPrice/config"
	"filmPrice/internal/apps"
	"filmPrice/internal/apps/task/dao"
	"filmPrice/internal/models"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/robfig/cron/v3"
)

func init() {
	// 启动定时任务扫描
	config.GetCron().AddFunc("@every 10m", SyncDBCronTask)
}

var FuncMap = map[string]func(params models.CustomMap) error{
	"taobaoSync": func(params models.CustomMap) error {

		return nil
	},
}

func cronAdd(task *dao.TaskModel) {
	fn, ok := FuncMap[task.FuncName]
	if !ok {
		apps.Log.Printf("Cron任务: %v 未找到函数: %v\n", task.Name, task.FuncName)
		return
	}

	entryID, err := config.GetCron().AddFunc(task.Cron, func() {
		taskLogID := TaskLogInit(&dao.TaskLogModel{
			TaskID: task.ID,
			Name:   fmt.Sprintf("[定时任务] %v", task.Name),
		})
		task.Params["log_id"] = taskLogID
		if err := fn(task.Params); err != nil {
			TaskLogUpdate(taskLogID, fmt.Sprintf("执行失败: %v", err))
			TaskLogStatusUpdate(taskLogID, "failed")
		}

		TaskLogStatusUpdate(taskLogID, "success")
		TaskLogUpdate(taskLogID, "执行完成")
	})
	if err != nil {
		apps.Log.Printf("Cron任务: %v 新增失败：%v\n", task.Name, err)
		return
	}

	// 更新数据库
	if err := config.GetDB().Model(task).Update("cron_id", int(entryID)).Error; err != nil {
		apps.Log.Printf("[Cron写入失败] %s: EntryID = %d 错误: %v", task.Name, entryID, err)
		return
	}

	apps.Log.Printf("[Cron注册成功] %s: EntryID = %d", task.Name, entryID)
}

func cronRemove(task *dao.TaskModel) {
	config.GetCron().Remove(cron.EntryID(task.CronID))
	apps.Log.Printf("Cron任务: %v EntryID: %v 删除成功\n", task.Name, task.CronID)
}

// SyncDBCronTask 同步数据库中的定时任务
func SyncDBCronTask() {
	var tasks []*dao.TaskModel

	if err := config.GetDB().Where("type = cron").Find(&tasks).Error; err != nil {
		apps.Log.Printf("Cron初始化错误：%v", err.Error())
	}

	entryIDSet := mapset.NewSet[int]()
	for _, e := range config.GetCron().Entries() {
		entryIDSet.Add(int(e.ID))
	}

	for _, task := range tasks {
		switch {
		case task.IsActive && task.CronID == 0:
			cronAdd(task)
		case !task.IsActive && task.CronID != 0 && entryIDSet.Contains(task.CronID):
			cronRemove(task)
		case task.IsActive && task.CronID != 0 && !entryIDSet.Contains(task.CronID):
			cronAdd(task)
		}
	}
}

// TaskRun 手动运行任务
func (s *service) TaskRun(id string) {
	var task *dao.TaskModel
	if err := s.db.Where("id = ?", id).First(&task).Error; err != nil {
		s.l.Printf("[TaskRun] 查询失败: %v ", err)
		return
	}

	taskLogID := TaskLogInit(&dao.TaskLogModel{
		TaskID: task.ID,
		Name:   fmt.Sprintf("[手动执行] %v", task.Name),
	})

	fn, ok := FuncMap[task.FuncName]
	if !ok {
		TaskLogUpdate(taskLogID, fmt.Sprintf("函数不存在: %v ", task.FuncName))
		TaskLogStatusUpdate(taskLogID, "error")
		return
	}

	TaskLogUpdate(taskLogID, fmt.Sprintf("开始执行: %v ", task.FuncName))
	task.Params["log_id"] = taskLogID
	if err := fn(task.Params); err != nil {
		TaskLogUpdate(taskLogID, fmt.Sprintf("执行失败: %v ", err))
		TaskLogStatusUpdate(taskLogID, "failed")
		return
	}

	TaskLogUpdate(taskLogID, fmt.Sprintf("执行成功: %v ", task.FuncName))
	TaskLogStatusUpdate(taskLogID, "success")
	return
}
