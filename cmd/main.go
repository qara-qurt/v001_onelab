package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"v001_onelab/configs"
	"v001_onelab/internal/repository"
	"v001_onelab/internal/service"
	rest "v001_onelab/internal/transport/http"
)

func main() {
	run()
}

func run() {
	config, err := configs.New()
	if err != nil {
		log.Fatal("cannot read config files")
	}

	repo := repository.New()
	service := service.New(repo)
	handler := rest.New(service)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.PORT),
		Handler: handler.InitRouter(),
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server started on %s \n", config.PORT)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server stopped gracefully")
}
