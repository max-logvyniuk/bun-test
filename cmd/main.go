package main

import (
	"context"
	"encoding/json"
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
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"

	"github.com/max-logvyniuk/bun-test/cmd/migrations"
	"github.com/max-logvyniuk/bun-test/lib/shared"
	"github.com/max-logvyniuk/bun-test/lib/svc"
	"github.com/max-logvyniuk/bun-test/lib/svc/app"
)

func main() {
	godotenv.Load()
	ctx := context.Background()

	app.Init(migrations.Migrations)

	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware(
			reqlog.FromEnv("BUNDEBUG"),
		)),
	)

	router.GET("/data", getMessageHandler(ctx, app.ApplicationService))
	router.POST("/data", createMessageHandler(ctx, app.ApplicationService))

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	address := host + ":" + port

	httpLn, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	handler := http.Handler(router)

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

func getMessageHandler(ctx context.Context, appSvc svc.Service) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		data, err := appSvc.GetAllData(ctx)

		if err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"data": data,
		})
	}
}

func createMessageHandler(ctx context.Context, appSvc svc.Service) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		dto := shared.DataCreate{}

		err := json.NewDecoder(req.Body).Decode(&dto)

		if err != nil {
			return err
		}

		data, err := appSvc.CreateData(ctx, dto)

		if err != nil {
			return err
		}

		return bunrouter.JSON(w, bunrouter.H{
			"data": data,
		})
	}
}
