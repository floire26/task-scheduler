package dto

type AddTaskRequest struct {
	Name    string `json:"name" binding:"required"`
	Minute  string `json:"minute" binding:"required"`
	Hour    string `json:"hour"  binding:"required"`
	Dom     string `json:"dom"  binding:"required"`
	Month   string `json:"month"  binding:"required"`
	Dow     string `json:"dow"  binding:"required"`
	FuncOpt uint8  `json:"func_opt"  binding:"required,min=1,max=2"`
}
