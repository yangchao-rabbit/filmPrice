package http

import (
	"filmPrice/internal/apps/film/model"
	"filmPrice/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FilmPriceHistoryList
//
// @Summary 胶卷历史价格列表
// @Description 获取胶卷历史价格链接列表
// @Tags FilmPriceHistory
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query model.FilmPriceHistoryListReq true "query"
// @Success 200 {object} models.Response{data=model.FilmPriceHistoryListResp} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /price-history [get]
func (h *handler) FilmPriceHistoryList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.FilmPriceHistoryListReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		list, err := h.svc.FilmPriceHistoryList(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(list))
	}
}

// FilmPriceHistoryDetail
//
// @Summary 获取胶卷历史价格详情
// @Description 获取胶卷历史价格详情
// @Tags FilmPriceHistory
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID"
// @Success 200 {object} models.Response{data=dao.FilmPriceHistoryModel} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /price-history/{id} [get]
func (h *handler) FilmPriceHistoryDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		cloud, err := h.svc.FilmPriceHistoryDetail(req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(cloud))
	}
}
