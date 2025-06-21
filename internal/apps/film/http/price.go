package http

import (
	"filmPrice/internal/apps/film/model"
	"filmPrice/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FilmPriceList
//
// @Summary 胶卷价格列表
// @Description 获取胶卷价格链接列表
// @Tags FilmPrice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query model.FilmPriceListReq true "query"
// @Success 200 {object} models.Response{data=model.FilmPriceListResp} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /price [get]
func (h *handler) FilmPriceList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.FilmPriceListReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		list, err := h.svc.FilmPriceList(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(list))
	}
}

// FilmPriceDetail
//
// @Summary 获取胶卷价格详情
// @Description 获取胶卷价格详情
// @Tags FilmPrice
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "胶卷价格ID"
// @Success 200 {object} models.Response{data=dao.FilmPriceModel} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /price/{id} [get]
func (h *handler) FilmPriceDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		cloud, err := h.svc.FilmPriceDetail(req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(cloud))
	}
}
