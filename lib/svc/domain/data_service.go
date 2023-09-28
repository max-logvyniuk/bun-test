package domain

import (
	"context"

	"github.com/max-logvyniuk/bun-test/lib/shared"
)

// DataService is a service which returns expected types of client instead of httpResponse.
type DataService interface {
	// GetAllData will return all data.
	GetAllData(context context.Context) ([]shared.Data, error)
	// Create new data and return it.
	Create(ctx context.Context, dto shared.DataCreate) (shared.Data, error)
}

type dataService struct {
	dataRepository DataRepository
}

// NewDataService creates a new instance of DataService.
func NewDataService(
	_ context.Context,
	dataRepository DataRepository,
) DataService {
	return &dataService{
		dataRepository: dataRepository,
	}
}

// GetAllData return list of all data.
func (svc *dataService) GetAllData(ctx context.Context) ([]shared.Data, error) {
	return svc.dataRepository.GetAllData(ctx)
}

// Create method create new instance of project.
func (svc *dataService) Create(ctx context.Context, dto shared.DataCreate) (shared.Data, error) {
	return svc.dataRepository.Create(ctx, dto)
}
