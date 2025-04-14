package router

import (
	"github.com/joaoasantana/e-product-service/app"
	"github.com/joaoasantana/e-product-service/internal/application/service"
	"github.com/joaoasantana/e-product-service/internal/infrastructure/store"
	"github.com/joaoasantana/e-product-service/internal/presentation/http/handler"
)

func NewProductRouter(startup *app.Startup) {
	dbConnection := startup.DBClient.Database(startup.Config.Database.Name)

	categoryRepository := store.NewMongoCategoryRepository(dbConnection)
	productRepository := store.NewMongoProductRepository(dbConnection)

	productService := service.NewProductService(categoryRepository, productRepository)
	productHandler := handler.NewProductHandler(productService)

	group := startup.Router.Group(startup.Config.BaseURL + "/products")
	{
		group.POST("/", productHandler.Create)

		group.GET("/", productHandler.FindAll)
		group.GET("/:id", productHandler.FindByID)
	}
}
