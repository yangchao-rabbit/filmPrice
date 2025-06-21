package http

import (
	"filmPrice/internal/apps/auth/model"
	"filmPrice/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LocalLogin
//
// @Summary 本地登录
// @Description 本地登录
// @Tags auth
// @Accept json
// @Produce json
// @Param req body model.LocalLoginReq true "请求参数"
// @Success 200 {object} models.Response{data=string} "成功"
// @Failure 400 {object} models.Response "失败"
// @Router /auth/local-login [post]
func (h *handler) LocalLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.LocalLoginReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		token, err := h.svc.LocalLogin(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		c.JSON(http.StatusOK, models.SuccessResp(token))
	}
}
