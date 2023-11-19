package shared

import "github.com/floire26/task-scheduler/model"

func ChooseOption(task model.Task, crontab string, opt uint8) (model.Task, error) {
	switch opt {
	case 1:
		task.Function = model.NewPrintMethod(task, crontab)
	case 2:
		task.Function = model.NewPrintWithTimeMethod(task, crontab)
	default:
		return task, ErrInvalidFuncOpt
	}

	return task, nil
}
