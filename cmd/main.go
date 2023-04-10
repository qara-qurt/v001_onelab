package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"v001_onelab/configs"
	"v001_onelab/internal/repository"
	"v001_onelab/internal/service"
	rest "v001_onelab/internal/transport/http"
	"v001_onelab/pkg/database/postgres"

	_ "v001_onelab/docs"
)

// @title Library
// @version 1.0
// @description REST API Library

// @contact.name Serikov Dias
// @contact.email serikov.2002.12@gmail.com

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	run()
}

func run() {
	config, err := configs.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := postgres.NewDatabasePSQL(config)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	service := service.New(repo, config)
	handler := rest.New(service)

	srv := handler.InitRouter()

	go func() {
		if err := srv.Start(fmt.Sprintf(":%s", config.PORT)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	gracefulShutdown(srv)

}

func gracefulShutdown(srv *echo.Echo) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server stopped gracefully")
}
