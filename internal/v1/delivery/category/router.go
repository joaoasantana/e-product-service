package category

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/app/category"
	store "github.com/joaoasantana/e-product-service/internal/v1/store/category"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(engine *gin.Engine, dbConn *mongo.Database) {
	categoryRepo := store.NewMongoRepository(dbConn)
	categoryService := category.NewService(categoryRepo)

	h := newHandler(categoryService)

	v1 := engine.Group("api/v1/categories")
	{
		v1.POST("/", h.createCategory)

		v1.GET("/", h.findAllCategories)
		v1.GET("/:id", h.findCategoryByID)
	}
}
