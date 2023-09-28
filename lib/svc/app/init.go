package app

import (
	"context"
	"sync"

	"github.com/max-logvyniuk/bun-test/lib/svc"
	"github.com/max-logvyniuk/bun-test/lib/svc/domain"
	dbclient "github.com/max-logvyniuk/bun-test/pkg/database/client"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var (
	ctx                context.Context
	database           *bun.DB
	onceInitAppService sync.Once
	// ApplicationService is a single instance of the application service.
	// Once Init was called, it will setup the application service.
	ApplicationService svc.Service
)

func Init(mgr *migrate.Migrations) {
	var err error

	onceInitAppService.Do(func() {
		ctx = context.Background()
		database, err = dbclient.NewDatabaseFactory().CreateConnection(ctx, mgr)
	})

	if err != nil {
		panic(err)
	}

	dataRepository := domain.NewDataRepository(database)
	dataService := domain.NewDataService(ctx, dataRepository)

	appService := NewApplicationService(dataService)

	ApplicationService = appService
}
