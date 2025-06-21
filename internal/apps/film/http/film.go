package http

import (
	"filmPrice/internal/apps/film/model"
	"filmPrice/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FilmList
//
// @Summary 胶卷列表
// @Description 获取胶卷列表
// @Tags Film
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query model.FilmListReq true "query"
// @Success 200 {object} models.Response{data=model.FilmListResp} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /film [get]
func (h *handler) FilmList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.FilmListReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		list, err := h.svc.FilmList(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(list))
	}
}

// FilmDetail
//
// @Summary 获取胶卷详情
// @Description 获取胶卷详情
// @Tags Film
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "胶卷ID"
// @Success 200 {object} models.Response{data=dao.FilmModel} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /film/{id} [get]
func (h *handler) FilmDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		cloud, err := h.svc.FilmDetail(req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(cloud))
	}
}

// FilmCreate
//
// @Summary 创建胶卷
// @Description 创建胶卷
// @Tags Film
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param param body model.FilmCreateReq true "入参"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /film [post]
func (h *handler) FilmCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.FilmCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.FilmCreate(&req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// FilmUpdate
//
// @Summary 更新胶卷
// @Description 更新胶卷
// @Tags Film
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "胶卷ID"
// @Param param body model.FilmUpdateReq true "胶卷信息"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /film/{id} [put]
func (h *handler) FilmUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req model.FilmUpdateReq
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

		if err := h.svc.FilmUpdate(uri.ID, &req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// FilmDelete
//
// @Summary      删除胶卷
// @Description  删除胶卷
// @Tags         Film
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id   path      string  true  "胶卷ID"
// @Success      200  {object}  models.Response "{"code": 0, "data": {}, "msg": "success"}"
// @Failure      400  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure      500  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router       /film/{id} [delete]
func (h *handler) FilmDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.FilmDelete(req.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}
