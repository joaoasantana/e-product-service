package category

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/app/category"
	"github.com/joaoasantana/e-product-service/internal/v1/infra/store"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(router *gin.RouterGroup, dbConn *mongo.Database) {
	categoryRepo := store.NewMongoCategoryRepository(dbConn)
	categoryService := category.NewService(categoryRepo)

	h := newHandler(categoryService)

	v1 := router.Group("/categories")
	{
		v1.POST("/", h.createCategory)

		v1.GET("/", h.findAllCategories)
		v1.GET("/:id", h.findCategoryByID)
	}
}
