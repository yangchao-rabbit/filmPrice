package http

import (
	"filmPrice/internal/apps/task/model"
	"filmPrice/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TaskList
//
// @Summary 任务列表
// @Description 获取任务列表
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param query query model.TaskListReq true "query"
// @Success 200 {object} models.Response{data=model.TaskListResp} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /task [get]
func (h *handler) TaskList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.TaskListReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		list, err := h.svc.TaskList(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(list))
	}
}

// TaskDetail
//
// @Summary 获取任务详情
// @Description 获取任务详情
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "任务ID"
// @Success 200 {object} models.Response{data=dao.TaskModel} "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /task/{id} [get]
func (h *handler) TaskDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		detail, err := h.svc.TaskGet(req.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(detail))
	}
}

// TaskCreate
//
// @Summary 创建任务
// @Description 创建任务
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param param body model.TaskCreateReq true "入参"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /task [post]
func (h *handler) TaskCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.TaskCreateReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.TaskCreate(&req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// TaskUpdate
//
// @Summary 更新任务
// @Description 更新任务
// @Tags Task
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param id path string true "任务ID"
// @Param param body model.TaskUpdateReq true "任务信息"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router /task/{id} [put]
func (h *handler) TaskUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req model.TaskUpdateReq
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

		if err := h.svc.TaskUpdate(uri.ID, &req); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// TaskDelete
//
// @Summary      删除任务
// @Description  删除任务
// @Tags         Task
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id   path      string  true  "任务ID"
// @Success      200  {object}  models.Response "{"code": 0, "data": {}, "msg": "success"}"
// @Failure      400  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Failure      500  {object}  models.Response "{"code": 1, "data": {}, "msg": "err"}"
// @Router       /task/{id} [delete]
func (h *handler) TaskDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		if err := h.svc.TaskDelete(req.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// TaskTestCron
//
// @Summary 测试cron
// @Description 测试cron
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param object body model.TaskTestCronReq true "cron表达式"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /task/test-cron [post]
func (h *handler) TaskTestCron() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.TaskTestCronReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		resp, err := h.svc.TestCron(req.Spec)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.FailResp(1, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, models.SuccessResp(resp))
	}
}

// TaskRun
//
// @Summary 运行任务
// @Description 运行任务
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param object body model.IDReq true "任务ID"
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Failure 400 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Failure 500 {object} models.Response "{"code": 1, "data": "", "msg": "error"}"
// @Router /task/run [post]
func (h *handler) TaskRun() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req model.IDReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, models.FailResp(1, err.Error()))
			return
		}

		go h.svc.TaskRun(req.ID)

		ctx.JSON(http.StatusOK, models.SuccessResp(nil))
	}
}

// TaskFuncList
//
// @Summary 获取任务列表
// @Description 任务列表
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Router /task/func [get]
func (h *handler) TaskFuncList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.SuccessResp(h.svc.TaskFuncList()))
	}
}

// TaskCurrentCron
//
// @Summary 获取当前任务列表
// @Description 当前任务列表
// @Tags Task
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response "{"code": 0, "data": {}, "msg": "ok"}"
// @Router /task/cur-cron [get]
func (h *handler) TaskCurrentCron() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.SuccessResp(h.svc.TaskCurrentCron()))
	}
}
