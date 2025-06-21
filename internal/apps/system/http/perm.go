package http

import (
	"filmPrice/internal/apps/system/model"
	"filmPrice/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PermList
//
// @Summary 权限列表
// @Description 权限列表
// @Tags System
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query model.PermListReq true "query"
// @Success 200 {object} models.Response{data=model.PermListResp} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /system/perm [get]
func (h *handler) PermList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.PermListReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		list, err := h.svc.PermList(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(list))
	}
}
