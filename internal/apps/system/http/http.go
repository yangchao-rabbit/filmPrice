package http

import (
	"filmPrice/internal/apps"
	"filmPrice/internal/apps/system"
	"filmPrice/internal/apps/system/service"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

func init() {
	apps.RegistryGin(&handler{
		svc: apps.GetImplSvc(system.AppName).(service.Service),
	})
}

type Handler interface {
	i()

	UserList() gin.HandlerFunc
	UserDetail() gin.HandlerFunc
	UserCreate() gin.HandlerFunc
	UserUpdate() gin.HandlerFunc
	UserDelete() gin.HandlerFunc

	GroupList() gin.HandlerFunc
	GroupDetail() gin.HandlerFunc
	GroupCreate() gin.HandlerFunc
	GroupUpdate() gin.HandlerFunc
	GroupDelete() gin.HandlerFunc

	PermList() gin.HandlerFunc
}

type handler struct {
	svc service.Service
}

func (*handler) i() {}

func (*handler) Name() string {
	return system.AppName
}

func (h *handler) Registry(r gin.IRouter) {
	userAPI := r.Group("/user")
	{
		userAPI.GET("", h.UserList())
		userAPI.GET("/:id", h.UserDetail())
		userAPI.POST("", h.UserCreate())
		userAPI.PUT("/:id", h.UserUpdate())
		userAPI.DELETE("/:id", h.UserDelete())
	}

	groupAPI := r.Group("/group")
	{
		groupAPI.GET("", h.GroupList())
		groupAPI.GET("/:id", h.GroupDetail())
		groupAPI.POST("", h.GroupCreate())
		groupAPI.PUT("/:id", h.GroupUpdate())
		groupAPI.DELETE("/:id", h.GroupDelete())
	}

	permAPI := r.Group("/perm")
	{
		permAPI.GET("", h.PermList())
	}
}
