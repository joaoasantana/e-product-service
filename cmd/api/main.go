package main

import (
	"context"
	"github.com/joaoasantana/e-product-service/app"
	"github.com/joaoasantana/e-product-service/internal/presentation/http/router"
	"time"
)

func main() {
	startup := app.NewStartup(app.LoadAppConfig())
	defer closeConnections(startup)

	router.NewCategoryRouter(startup)
	router.NewProductRouter(startup)

	if err := startup.Router.Run(startup.Config.Server.Port); err != nil {
		panic(err)
	}
}

func closeConnections(startup *app.Startup) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := startup.DBClient.Disconnect(ctx); err != nil {
		panic(err)
	}
	if err := startup.ZapLogger.Sync(); err != nil {
		panic(err)
	}
}
