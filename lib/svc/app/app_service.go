package app

import (
	"context"

	"github.com/max-logvyniuk/bun-test/lib/shared"
	"github.com/max-logvyniuk/bun-test/lib/svc"
	"github.com/max-logvyniuk/bun-test/lib/svc/domain"
)

type appService struct {
	dataService domain.DataService
}

// NewApplicationService will instantiate a new application service.
func NewApplicationService(
	dataService domain.DataService,
) svc.Service {
	return &appService{
		dataService: dataService,
	}
}

// CreateData create new data object.
func (svc *appService) CreateData(ctx context.Context, dto shared.DataCreate) (shared.Data, error) {
	return svc.dataService.Create(ctx, dto)
}

// GetAllData query all data objects.
func (svc *appService) GetAllData(ctx context.Context) ([]shared.Data, error) {
	return svc.dataService.GetAllData(ctx)
}
