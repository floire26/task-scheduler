package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/floire26/task-scheduler/api/handler"
	"github.com/floire26/task-scheduler/api/middleware"
	"github.com/floire26/task-scheduler/api/usecase"
	"github.com/floire26/task-scheduler/shared"
	"github.com/gin-gonic/gin"
)

func Init() {
	cfg := shared.LoadConfig("../.env")
	router := gin.New()
	taskUc := usecase.NewTaskUsecase()
	taskUc.StartScheduler()
	taskHandler := handler.NewTaskHandler(taskUc)
	router.Use(middleware.ErrorResponse())
	router.POST("/tasks", taskHandler.HandleAddTaskToScheduler)
	router.DELETE("/tasks/:id", taskHandler.HandleRemoveTaskFromScheduler)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	taskUc.StopScheduler()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.TimeoutDur)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
