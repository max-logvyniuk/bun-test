package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// "github.com/klauspost/compress/gzhttp"
	// "github.com/rs/cors"

	"github.com/joho/godotenv"

	"github.com/max-logvyniuk/bun-test/cmd/migrations"
	"github.com/max-logvyniuk/bun-test/lib/svc/app"
	"github.com/max-logvyniuk/bun-test/lib/svc/transport"
)

func main() {
	godotenv.Load()
	ctx := context.Background()

	app.Init(migrations.Migrations)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	address := host + ":" + port

	httpLn, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	handler := transport.NewHTTPTransport(ctx, app.ApplicationService)

	httpServer := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handler,
	}

	fmt.Println("listening on " + address)
	go func() {
		if err := httpServer.Serve(httpLn); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Press CTRL+C to exit...")
	fmt.Println(waitExitSignal())

	// Graceful shutdown.
	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
}

func waitExitSignal() os.Signal {
	ch := make(chan os.Signal, 3)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-ch
}
