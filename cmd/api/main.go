package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xenakil/cheddartrack/internal/config"
	"github.com/xenakil/cheddartrack/internal/db"
	"github.com/xenakil/cheddartrack/internal/handler"
	"github.com/xenakil/cheddartrack/internal/repository"
	"github.com/xenakil/cheddartrack/internal/service"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    _ "github.com/xenakil/cheddartrack/docs"
)

func main() {
	cfg := config.LoadConfig()

	// Connecting to database
	db.Init(cfg)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	txnRepo := repository.NewTransactionRepository(db.DB)
	txnSvc := service.NewTransactionService(txnRepo)
	txnHandler := handler.NewTransactionHanlder(txnSvc)
	txnHandler.RegisterRouter(r)
	
	r.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"status": "ok", "env": cfg.Env})
	})

	srv := &http.Server{
		Addr: ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
    fmt.Println("Server running on http://localhost:%s\n" + cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
    fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
	}

    fmt.Println("Server exiting")
}