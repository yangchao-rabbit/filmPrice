package http

import (
	"filmPrice/internal/apps/system/model"
	"filmPrice/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserList
//
// @Summary 用户列表
// @Description 用户列表
// @Tags System
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query model.UserListReq true "query"
// @Success 200 {object} models.Response{data=model.UserListResp} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /system/user [get]
func (h *handler) UserList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.UserListReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		list, err := h.svc.UserList(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(list))
	}
}

// UserDetail
//
// @Summary 用户详情
// @Description 用户详情
// @Tags System
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "用户ID"
// @Success 200 {object} models.Response{data=dao.SystemUserModel} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /system/user/{id} [get]
func (h *handler) UserDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		cloud, err := h.svc.UserGet(req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(cloud))
	}
}

// UserCreate
//
// @Summary 创建用户
// @Description 创建用户
// @Tags System
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param param body model.UserCreateReq true "入参"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /system/user [post]
func (h *handler) UserCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.UserCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.UserCreate(&req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// UserUpdate
//
// @Summary 用户更新
// @Description 用户更新
// @Tags System
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "用户ID"
// @Param param body model.UserUpdateReq true "用户信息"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /system/user/{id} [put]
func (h *handler) UserUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req model.UserUpdateReq
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

		if err := h.svc.UserUpdate(uri.ID, &req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// UserDelete
//
// @Summary      删除用户
// @Description  删除用户
// @Tags         System
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id   path      string  true  "用户ID"
// @Success      200  {object}  models.Response "{"code": 0, "data": {}, "msg": "success"}"
// @Failure      400  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure      500  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router       /system/user/{id} [delete]
func (h *handler) UserDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.UserDelete(req.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}
