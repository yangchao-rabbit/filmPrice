package http

import (
	"filmPrice/internal/apps/film/model"
	"filmPrice/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FilmLinkList
//
// @Summary 胶卷链接列表
// @Description 获取胶卷链接列表
// @Tags FilmLink
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query model.FilmLinkListReq true "query"
// @Success 200 {object} models.Response{data=model.FilmLinkListResp} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /link [get]
func (h *handler) FilmLinkList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.FilmLinkListReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		list, err := h.svc.FilmLinkList(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(list))
	}
}

// FilmLinkDetail
//
// @Summary 获取胶卷链接详情
// @Description 获取胶卷链接详情
// @Tags FilmLink
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "胶卷链接ID"
// @Success 200 {object} models.Response{data=dao.FilmLinkModel} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /link/{id} [get]
func (h *handler) FilmLinkDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		cloud, err := h.svc.FilmLinkDetail(req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(cloud))
	}
}

// FilmLinkCreate
//
// @Summary 创建胶卷链接
// @Description 创建胶卷链接
// @Tags FilmLink
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param param body model.FilmLinkCreateReq true "入参"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /link [post]
func (h *handler) FilmLinkCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.FilmLinkCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.FilmLinkCreate(&req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// FilmLinkUpdate
//
// @Summary 更新胶卷链接
// @Description 更新胶卷链接
// @Tags FilmLink
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "胶卷链接ID"
// @Param param body model.FilmLinkUpdateReq true "胶卷链接信息"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /link/{id} [put]
func (h *handler) FilmLinkUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req model.FilmLinkUpdateReq
			uri model.IDReq
		)
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := ctx.ShouldBindUri(&uri); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.FilmLinkUpdate(uri.ID, &req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// FilmLinkDelete
//
// @Summary      删除胶卷链接
// @Description  删除胶卷链接
// @Tags         FilmLink
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id   path      string  true  "胶卷链接ID"
// @Success      200  {object}  models.Response "{"code": 0, "data": {}, "msg": "success"}"
// @Failure      400  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure      500  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router       /link/{id} [delete]
func (h *handler) FilmLinkDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.FilmLinkDelete(req.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}
