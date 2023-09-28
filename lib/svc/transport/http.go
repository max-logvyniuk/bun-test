package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/max-logvyniuk/bun-test/lib/shared"
	"github.com/max-logvyniuk/bun-test/lib/svc"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func NewHTTPTransport(ctx context.Context, service svc.Service) http.Handler {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware(
			reqlog.FromEnv("BUNDEBUG"),
		)),
	)

	router.GET("/data", getMessageHandler(ctx, service))
	router.POST("/data", createMessageHandler(ctx, service))

	handler := http.Handler(router)

	return handler
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
