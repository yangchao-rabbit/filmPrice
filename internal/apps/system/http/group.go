package http

import (
	"filmPrice/internal/apps/system/model"
	"filmPrice/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GroupList
//
// @Summary 组列表
// @Description 组列表
// @Tags System
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query model.GroupListReq true "query"
// @Success 200 {object} models.Response{data=model.GroupListResp} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /system/group [get]
func (h *handler) GroupList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.GroupListReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		list, err := h.svc.GroupList(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(list))
	}
}

// GroupDetail
//
// @Summary 组详情
// @Description 组详情
// @Tags System
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "组ID"
// @Success 200 {object} models.Response{data=dao.SystemGroupModel} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /system/group/{id} [get]
func (h *handler) GroupDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		cloud, err := h.svc.GroupGet(req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(cloud))
	}
}

// GroupCreate
//
// @Summary 创建组
// @Description 创建组
// @Tags System
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param param body model.GroupCreateReq true "入参"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /system/group [post]
func (h *handler) GroupCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.GroupCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.GroupCreate(&req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// GroupUpdate
//
// @Summary 组更新
// @Description 组更新
// @Tags System
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "组ID"
// @Param param body model.GroupUpdateReq true "组信息"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /system/group/{id} [put]
func (h *handler) GroupUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req model.GroupUpdateReq
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

		if err := h.svc.GroupUpdate(uri.ID, &req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// GroupDelete
//
// @Summary      删除组
// @Description  删除组
// @Tags         System
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id   path      string  true  "组ID"
// @Success      200  {object}  models.Response "{"code": 0, "data": {}, "msg": "success"}"
// @Failure      400  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure      500  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router       /system/group/{id} [delete]
func (h *handler) GroupDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.GroupDelete(req.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}
