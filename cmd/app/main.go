// Package main provides the main entry point for the  CRUD API
//
//	@title						CRUD API
//	@version					1.2.0
//	@description				CRUD Api for employees
//	@termsOfService				http://swagger.io/terms/
//
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//
//	@host						localhost:8080
//	@BasePath					/api/v1
//
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

	"github.com/MaulanaAhmadSulami/juke_test.git/internal/config"
	"github.com/MaulanaAhmadSulami/juke_test.git/internal/db"
	employeeRepo "github.com/MaulanaAhmadSulami/juke_test.git/internal/repository/postgres/employee"
	employeeService "github.com/MaulanaAhmadSulami/juke_test.git/internal/service/employee"
	employeeHandler "github.com/MaulanaAhmadSulami/juke_test.git/internal/server/http/handler/employee"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	_ "github.com/MaulanaAhmadSulami/juke_test.git/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const VERSION = "1.1.4"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to create logger:", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infow("start api", "version", VERSION)

	cfg, err := config.Load()
	if err != nil {
		sugar.Fatalw("Failed to load config", "error", err)
	}

	database, err := db.NewPostgresDB(cfg.GetDBConnectionString(), cfg.DB)
	if err != nil {
		sugar.Fatalw("feailed to connect to db", err)
	}
	defer database.Close()

	sugar.Info("db connected")


	empRepo := employeeRepo.NewEmployeeStore(database)
	empService := employeeService.NewEmployeeService(empRepo)

	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"OK","version":"` + VERSION + `"}`))
	})

	router.Get("/swagger/*", httpSwagger.WrapHandler)
	router.Route("/api/v1/employees", employeeHandler.RegisterRoute(empService, sugar))

	sugar.Info("Routes registered")

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	go func() {
		sugar.Infow("Server started", "address", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sugar.Fatalw("Server failed", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	sugar.Info("shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		sugar.Fatalw("server forced shutdown", "error", err)
	}

	sugar.Info("server stop")
}