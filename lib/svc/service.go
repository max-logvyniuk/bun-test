package svc

import (
	"context"

	"github.com/max-logvyniuk/bun-test/lib/shared"
)

// Service defines all the APIs available for the service.
type Service interface {
	// CreateData create new data.
	CreateData(ctx context.Context, dto shared.DataCreate) (shared.Data, error)
	// GetAllData return all data.
	GetAllData(ctx context.Context) ([]shared.Data, error)
}
