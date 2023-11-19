package usecase

import (
	"fmt"
	"time"

	"github.com/adhocore/gronx"
	"github.com/aptible/supercronic/cronexpr"
	"github.com/floire26/task-scheduler/dto"
	"github.com/floire26/task-scheduler/model"
	"github.com/floire26/task-scheduler/shared"
	"github.com/robfig/cron/v3"
)

var (
	gron = gronx.New()
)

type taskUsecase struct {
	scheduler *cron.Cron
}

type TaskUsecase interface {
	StartScheduler()
	StopScheduler()
	AddTaskToScheduler(reqBody dto.AddTaskRequest) (cron.EntryID, error)
	RemoveTaskFromScheduler(id cron.EntryID) (cron.EntryID, error)
}

func NewTaskUsecase() TaskUsecase {
	return &taskUsecase{scheduler: cron.New()}
}

func (u *taskUsecase) StartScheduler() {
	u.scheduler.Start()
}

func (u *taskUsecase) StopScheduler() {
	u.scheduler.Stop()
}

func (u *taskUsecase) AddTaskToScheduler(reqBody dto.AddTaskRequest) (cron.EntryID, error) {
	task := model.Task{
		Name: reqBody.Name,
	}

	crontab := fmt.Sprintf("%s %s %s %s %s",
		reqBody.Minute,
		reqBody.Hour,
		reqBody.Dom,
		reqBody.Month,
		reqBody.Dow,
	)

	if !gron.IsValid(crontab) {
		return 0, shared.ErrInvalidCrontab
	}

	task, err := shared.ChooseOption(task, crontab, reqBody.FuncOpt)

	if err != nil {
		return 0, err
	}

	task.NextExecTime = cronexpr.MustParse(crontab).Next(time.Now()).Unix()
	res, err := u.scheduler.AddFunc(crontab, task.Function.Exec())

	if err != nil {
		return 0, shared.ErrAddTaskFailed
	}

	return res, err
}

func (u *taskUsecase) RemoveTaskFromScheduler(id cron.EntryID) (cron.EntryID, error) {
	var taskExist bool

	for _, sch := range u.scheduler.Entries() {
		if sch.ID == id {
			taskExist = true
			break
		}
	}

	if !taskExist {
		return id, shared.ErrTaskNotFound
	}

	u.scheduler.Remove(id)
	return id, nil
}
