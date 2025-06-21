package http

import (
	"filmPrice/internal/apps"
	"filmPrice/internal/apps/film"
	"filmPrice/internal/apps/film/service"
	"github.com/gin-gonic/gin"
)

var _ Handler = (*handler)(nil)

func init() {
	apps.RegistryGin(&handler{
		svc: apps.GetImplSvc(film.AppName).(service.Service),
	})
}

type Handler interface {
	i()

	FilmList() gin.HandlerFunc
	FilmDetail() gin.HandlerFunc
	FilmCreate() gin.HandlerFunc
	FilmUpdate() gin.HandlerFunc
	FilmDelete() gin.HandlerFunc

	FilmLinkList() gin.HandlerFunc
	FilmLinkDetail() gin.HandlerFunc
	FilmLinkCreate() gin.HandlerFunc
	FilmLinkUpdate() gin.HandlerFunc
	FilmLinkDelete() gin.HandlerFunc

	FilmPriceList() gin.HandlerFunc
	FilmPriceDetail() gin.HandlerFunc
}

type handler struct {
	svc service.Service
}

func (*handler) i() {}

func (*handler) Name() string {
	return film.AppName
}

func (h *handler) Registry(r gin.IRouter) {
	// 胶卷
	r.GET("", h.FilmList())
	r.GET("/:id", h.FilmDetail())
	r.POST("", h.FilmCreate())
	r.PUT("/:id", h.FilmUpdate())
	r.DELETE("/:id", h.FilmDelete())

	// 胶卷链接
	linkAPI := r.Group("/link")
	{
		linkAPI.GET("", h.FilmLinkList())
		linkAPI.GET("/:id", h.FilmLinkDetail())
		linkAPI.POST("", h.FilmLinkCreate())
		linkAPI.PUT("/:id", h.FilmLinkUpdate())
		linkAPI.DELETE("/:id", h.FilmLinkDelete())
	}

	// 胶卷价格
	priceAPI := r.Group("/price")
	{
		priceAPI.GET("", h.FilmPriceList())
		priceAPI.GET("/:id", h.FilmPriceDetail())
	}
}
