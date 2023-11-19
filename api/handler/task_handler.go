package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/floire26/task-scheduler/api/usecase"
	"github.com/floire26/task-scheduler/dto"
	"github.com/floire26/task-scheduler/shared"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type TaskHandler struct {
	taskUc usecase.TaskUsecase
}

func NewTaskHandler(taskUc usecase.TaskUsecase) TaskHandler {
	return TaskHandler{taskUc: taskUc}
}

func (h TaskHandler) HandleAddTaskToScheduler(c *gin.Context) {
	var reqBody dto.AddTaskRequest

	err := c.ShouldBindJSON(&reqBody)

	if err != nil {
		c.Error(err)
		return
	}

	resBody, err := h.taskUc.AddTaskToScheduler(reqBody)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("task with ID #%d has been created", resBody)})
}

func (h TaskHandler) HandleRemoveTaskFromScheduler(c *gin.Context) {
	str, match := c.Params.Get("id")

	if !match {
		c.Error(shared.ErrMissingID)
		return
	}

	id, err := strconv.Atoi(str)

	if err != nil {
		c.Error(shared.ErrInvalidIDType)
		return
	}

	resBody, err := h.taskUc.RemoveTaskFromScheduler(cron.EntryID(id))

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("task with ID #%d has been deleted", resBody)})
}
