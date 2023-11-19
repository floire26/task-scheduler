package model

import (
	"fmt"
	"time"

	"github.com/aptible/supercronic/cronexpr"
)

type functionOption interface {
	Exec() func()
}

type print struct {
	task    Task
	crontab string
}

type printWithTime struct {
	task    Task
	crontab string
}

func (fo print) Exec() func() {
	return func() {
		fo.task.NextExecTime = cronexpr.MustParse(fo.crontab).Next(time.Now()).Unix()
		fmt.Println(fo.task.Name, "has been executed")
	}
}

var (
	layoutFormat = "2006-01-02 15:04:05"
)

func (fo printWithTime) Exec() func() {
	return func() {
		nextTime := cronexpr.MustParse(fo.crontab).Next(time.Now())
		fo.task.NextExecTime = nextTime.Unix()
		fmt.Println(fo.task.Name, "has been executed at:", time.Now().UTC().Format(layoutFormat), ", next execution at", nextTime.UTC().Format(layoutFormat))
	}
}

func NewPrintMethod(task Task, crontab string) print {
	return print{
		task:    task,
		crontab: crontab,
	}
}

func NewPrintWithTimeMethod(task Task, crontab string) printWithTime {
	return printWithTime{
		task:    task,
		crontab: crontab}
}
