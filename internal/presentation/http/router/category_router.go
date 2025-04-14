package router

import (
	"github.com/joaoasantana/e-product-service/app"
	"github.com/joaoasantana/e-product-service/internal/application/service"
	"github.com/joaoasantana/e-product-service/internal/infrastructure/store"
	"github.com/joaoasantana/e-product-service/internal/presentation/http/handler"
)

func NewCategoryRouter(startup *app.Startup) {
	dbConnection := startup.DBClient.Database(startup.Config.Database.Name)

	categoryRepository := store.NewMongoCategoryRepository(dbConnection)

	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	group := startup.Router.Group(startup.Config.BaseURL + "/categories")
	{
		group.POST("/", categoryHandler.Create)

		group.GET("/", categoryHandler.FindAll)
		group.GET("/:id", categoryHandler.FindByID)
	}
}
