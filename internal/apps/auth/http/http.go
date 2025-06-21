package http

import (
	"filmPrice/internal/apps"
	"filmPrice/internal/apps/auth"
	"filmPrice/internal/apps/auth/service"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

func init() {
	apps.RegistryGin(&handler{
		svc: apps.GetImplSvc(auth.AppName).(service.Service),
	})
}

type Handler interface {
	i()

	LocalLogin() gin.HandlerFunc
}

type handler struct {
	svc service.Service
}

func (*handler) i() {}

func (*handler) Name() string {
	return auth.AppName
}

func (h *handler) Registry(r gin.IRouter) {
	r.POST("/local-login", h.LocalLogin())
}
