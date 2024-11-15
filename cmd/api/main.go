package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	// "simulation/internal/database/seeder"
	"simulation/internal/database/seeder"
	"simulation/internal/server"
	"strconv"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func gracefulShutdown(fiberServer *server.FiberServer) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := fiberServer.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
}

func main() {

	server := server.New()

	go func() {
		if err := seeder.SeedDatabase(server.DB.GetDB()); err != nil {
			log.Printf("Seeder Error: %v", err) // Gunakan log.Printf untuk melanjutkan eksekusi
		}
	}()

	server.RegisterFiberRoutes()
	go func() {
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		err := server.Listen(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(fmt.Sprintf("http server error: %s", err))
		}
	}()

	gracefulShutdown(server)
}
