package http

import (
	"filmPrice/internal/apps"
	"filmPrice/internal/apps/task"
	"filmPrice/internal/apps/task/service"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

func init() {
	apps.RegistryGin(&handler{
		svc: apps.GetImplSvc(task.AppName).(service.Service),
	})
}

type Handler interface {
	i()

	TaskList() gin.HandlerFunc
	TaskDetail() gin.HandlerFunc
	TaskCreate() gin.HandlerFunc
	TaskUpdate() gin.HandlerFunc
	TaskDelete() gin.HandlerFunc
	TaskTestCron() gin.HandlerFunc
	TaskRun() gin.HandlerFunc

	TaskFuncList() gin.HandlerFunc
	TaskCurrentCron() gin.HandlerFunc
}

type handler struct {
	svc service.Service
}

func (*handler) i() {}

func (*handler) Name() string {
	return task.AppName
}

func (h *handler) Registry(r gin.IRouter) {
	r.GET("", h.TaskList())
	r.GET("/:id", h.TaskDetail())
	r.POST("", h.TaskCreate())
	r.PUT("/:id", h.TaskUpdate())
	r.DELETE("/:id", h.TaskDelete())
	r.POST("/test-cron", h.TaskTestCron())
	r.POST("/run", h.TaskRun())
	r.GET("/func", h.TaskFuncList())
	r.GET("/cur-cron", h.TaskCurrentCron())
}
