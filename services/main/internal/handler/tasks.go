package handler

import (
	"github.com/gin-gonic/gin"
	taskspb "main-service/proto"
	"net/http"
)

type CreateTaskRequest struct {
	Name string `json:"name"`
	Body string `json:"body"`
}

type CreateTaskResponse struct {
	TaskId uint64 `json:"task_id"`
	Status string `json:"status"`
}

// @Summary CreateTask
// @Tags tasks
// @Description create task
// @ID login
// @Accept  json
// @Produce  json
// @Param input body CreateTaskRequest true "params"
// @Success 200 {object} CreateTaskResponse
// @Failure 400,401,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/create-task [post]
func (h *Handler) createTask(c *gin.Context) {
	username, err := h.getUsername(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}
	body := &CreateTaskRequest{}
	err = c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := h.tasks.CreateTask(&taskspb.CreateTaskRequest{
		Common: &taskspb.CommonRequest{
			Username: username,
		},
		Name: body.Name,
		Body: body.Body,
	})
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	statusCode := http.StatusOK
	if resp.Status != "ok" {
		statusCode = http.StatusBadRequest
	}
	c.JSON(statusCode, &CreateTaskResponse{
		TaskId: resp.TaskId,
		Status: resp.Status,
	})
}

type UpdateTaskRequest struct {
	TaskId uint64 `json:"task_id"`
	Name   string `json:"name"`
	Body   string `json:"body"`
}

type UpdateTaskResponse struct {
	Status string `json:"status"`
}

// @Summary UpdateTask
// @Tags tasks
// @Description update task
// @ID login
// @Accept  json
// @Produce  json
// @Param input body UpdateTaskRequest true "params"
// @Success 200 {object}  UpdateTaskResponse
// @Failure 400,401,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/create-task [post]
func (h *Handler) updateTask(c *gin.Context) {
	username, err := h.getUsername(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}
	body := &UpdateTaskRequest{}
	err = c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := h.tasks.UpdateTask(&taskspb.UpdateTaskRequest{
		Common: &taskspb.CommonRequest{
			Username: username,
		},
		TaskId: body.TaskId,
		Name:   body.Name,
		Body:   body.Body,
	})
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	statusCode := http.StatusOK
	if resp.Status != "ok" {
		statusCode = http.StatusBadRequest
	}
	c.JSON(statusCode, &UpdateTaskResponse{
		Status: resp.Status,
	})
}

type DeleteTaskRequest struct {
	TaskId uint64 `json:"task_id"`
}

type DeleteTaskResponse struct {
	Status string `json:"status"`
}

// @Summary DeleteTask
// @Tags tasks
// @Description delete task
// @ID login
// @Accept  json
// @Produce  json
// @Param input body DeleteTaskRequest true "params"
// @Success 200 {object}  DeleteTaskResponse
// @Failure 400,401,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/delete-task [post]
func (h *Handler) deleteTask(c *gin.Context) {
	username, err := h.getUsername(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}
	body := &DeleteTaskRequest{}
	err = c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := h.tasks.DeleteTask(&taskspb.DeleteTaskRequest{
		Common: &taskspb.CommonRequest{
			Username: username,
		},
		TaskId: body.TaskId,
	})
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	statusCode := http.StatusOK
	if resp.Status != "ok" {
		statusCode = http.StatusBadRequest
	}
	c.JSON(statusCode, &DeleteTaskResponse{
		Status: resp.Status,
	})
}

type GetTaskRequest struct {
	TaskId uint64 `json:"task_id"`
}

type GetTaskResponse struct {
	TaskId uint64 `json:"task_id"`
	Author string `json:"author"`
	Name   string `json:"name"`
	Body   string `json:"body"`
}

// @Summary GetTask
// @Tags tasks
// @Description get task
// @ID login
// @Accept  json
// @Produce  json
// @Param input body GetTaskRequest true "params"
// @Success 200 {object} GetTaskResponse
// @Failure 400,401,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/get-task [get]
func (h *Handler) getTask(c *gin.Context) {
	username, err := h.getUsername(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}
	body := &GetTaskRequest{}
	err = c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := h.tasks.GetTask(&taskspb.GetTaskRequest{
		Common: &taskspb.CommonRequest{
			Username: username,
		},
		TaskId: body.TaskId,
	})
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	if resp.Status != "ok" {
		newErrorResponse(c, resp.Status, http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, &GetTaskResponse{
		Author: resp.Task.Spec.Author,
		Name:   resp.Task.Spec.Name,
		Body:   resp.Task.Spec.Body,
		TaskId: resp.Task.Id,
	})
}

type GetListTasksRequest struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

type GetListTasksResponse struct {
	Tasks []*GetTaskResponse `json:"tasks"`
}

// @Summary GetListTasks
// @Tags tasks
// @Description get list of tasks
// @ID login
// @Accept  json
// @Produce  json
// @Param input body GetListTasksRequest true "params"
// @Success 200 {object} GetListTasksResponse
// @Failure 400,401,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/get-list-tasks [get]
func (h *Handler) getListTasks(c *gin.Context) {
	username, err := h.getUsername(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusUnauthorized)
		return
	}
	body := &GetListTasksRequest{}
	err = c.BindJSON(body)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := h.tasks.GetListTasks(&taskspb.GetListTasksRequest{
		Common: &taskspb.CommonRequest{
			Username: username,
		},
		Offset: body.Offset,
		Limit:  body.Limit,
	})
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	var tasks []*GetTaskResponse
	for _, task := range resp {
		tasks = append(tasks, &GetTaskResponse{
			Author: task.Task.Spec.Author,
			Name:   task.Task.Spec.Name,
			Body:   task.Task.Spec.Body,
			TaskId: task.Task.Id,
		})
	}
	c.JSON(http.StatusOK, &GetListTasksResponse{
		Tasks: tasks,
	})
}
