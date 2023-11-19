package shared

import "net/http"

type CustomError struct {
	Message string
	Code    int
}

func (e CustomError) Error() string {
	return e.Message
}

var (
	ErrMissingID      = &CustomError{"missing id in query parameter", http.StatusBadRequest}
	ErrInvalidIDType  = &CustomError{"invalid id type in parameter", http.StatusBadRequest}
	ErrEmptyName      = &CustomError{"name field must be filled", http.StatusBadRequest}
	ErrInvalidCrontab = &CustomError{"invalid crontab expression", http.StatusBadRequest}
	ErrInvalidFuncOpt = &CustomError{"function option must be between 1 or 2", http.StatusBadRequest}
	ErrAddTaskFailed  = &CustomError{"adding task to the scheduler failed", http.StatusBadRequest}
	ErrTaskNotFound   = &CustomError{"the specified task id is not found in the scheduler", http.StatusBadRequest}
)
