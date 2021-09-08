package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/artofimagination/ip-resolver/initialization"
	"github.com/artofimagination/ip-resolver/rest"

	"github.com/pkg/errors"
)

func main() {
	cfg := &initialization.Config{}
	initialization.InitConfig(cfg)

	r, err := rest.CreateRouting()
	if err != nil {
		log.Fatalf("Failed to create routing. %s\n", errors.WithStack(err))
	}

	// Create Server and Route Handlers
	//port := fmt.Sprintf(":%d", cfg.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting down")
	os.Exit(0)
}
