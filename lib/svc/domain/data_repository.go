package domain

import (
	"context"

	"github.com/max-logvyniuk/bun-test/lib/shared"
	"github.com/uptrace/bun"
)

// DataRepository is interface for Project repository.
type DataRepository interface {
	// GetAllData return all data array.
	GetAllData(ctx context.Context) ([]shared.Data, error)
	// Create new item of Data in DB.
	Create(ctx context.Context, p shared.DataCreate) (shared.Data, error)
}

type dataRepository struct {
	db *bun.DB
}

// NewDataRepository return instance of DataRepository.
func NewDataRepository(database *bun.DB) DataRepository {
	return &dataRepository{
		db: database,
	}
}

// GetAllData return all data array.
func (dr *dataRepository) GetAllData(ctx context.Context) ([]shared.Data, error) {
	data := make([]shared.Data, 0)
	if err := dr.db.NewSelect().Model(&data).OrderExpr("id ASC").Scan(ctx); err != nil {
		return data, err
	}

	return data, nil
}

// Create new item of Data in DB.
func (dr *dataRepository) Create(ctx context.Context, dc shared.DataCreate) (shared.Data, error) {
	d := shared.Data{
		Message: dc.Message,
	}

	_, err := dr.db.NewInsert().Model(&d).Exec(ctx)

	if err != nil {
		return shared.Data{}, err
	}

	data := shared.Data{
		Message: dc.Message,
	}

	return data, nil
}
