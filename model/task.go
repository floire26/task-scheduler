package model

import (
	"github.com/robfig/cron/v3"
)

type Task struct {
	ID           cron.EntryID
	Name         string
	NextExecTime int64
	Function     functionOption
}
