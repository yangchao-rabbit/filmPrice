package model

import (
	"filmPrice/internal/apps/film/dao"
	"filmPrice/internal/models"
)

type FilmListReq struct {
	models.ListBaseReq
	Brand  string `json:"brand" form:"brand"`
	ISO    string `json:"iso" form:"iso"`
	Type   string `json:"type" form:"type"`
	Format string `json:"format" form:"format"`
}

type FilmListResp struct {
	Total int64            `json:"total"`
	Rows  []*dao.FilmModel `json:"rows"`
}

type FilmCreateReq struct {
	Alias  string `json:"alias" form:"alias" binding:"required"`
	Name   string `json:"name" form:"name" binding:"required"`
	Brand  string `json:"brand" form:"brand" binding:"required"`
	ISO    string `json:"iso" form:"iso" binding:"required"`
	Type   string `json:"type" form:"type" binding:"required"`
	Format string `json:"format" form:"format" binding:"required"`
	Image  string `json:"image" form:"image"`
	Desc   string `json:"desc" form:"desc"`
}

type FilmUpdateReq struct {
	Alias  string `json:"alias" form:"alias"`
	Name   string `json:"name" form:"name"`
	Brand  string `json:"brand" form:"brand"`
	ISO    string `json:"iso" form:"iso"`
	Type   string `json:"type" form:"type"`
	Format string `json:"format" form:"format"`
	Image  string `json:"image" form:"image"`
	Desc   string `json:"desc" form:"desc"`
}

type IDReq struct {
	ID string `json:"id" form:"id" uri:"id" binding:"required"`
}

type FilmLinkListReq struct {
	models.ListBaseReq
}

type FilmLinkListResp struct {
	Total int64                `json:"total"`
	Rows  []*dao.FilmLinkModel `json:"rows"`
}

type FilmLinkCreateReq struct {
	FilmID   string `json:"film_id" form:"film_id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Platform string `json:"platform" form:"platform" binding:"required"`
	Url      string `json:"url" form:"url" binding:"required"`
	IsActive bool   `json:"is_active" form:"is_active"`
	Desc     string `json:"desc" form:"desc"`
}

type FilmLinkUpdateReq struct {
	FilmID   string `json:"film_id" form:"film_id"`
	Name     string `json:"name" form:"name" `
	Platform string `json:"platform" form:"platform" `
	Url      string `json:"url" form:"url" `
	IsActive bool   `json:"is_active" form:"is_active"`
	Desc     string `json:"desc" form:"desc"`
}

type FilmPriceListReq struct {
	models.ListBaseReq
}

type FilmPriceListResp struct {
	Total int64                 `json:"total"`
	Rows  []*dao.FilmPriceModel `json:"rows"`
}

type FilmPriceHistoryListReq struct {
	models.ListBaseReq
	LinkID string `json:"link_id" form:"link_id"`
}

type FilmPriceHistoryListResp struct {
	Total int64                        `json:"total"`
	Rows  []*dao.FilmPriceHistoryModel `json:"rows"`
}
