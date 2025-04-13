package product

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/app/product"
	"github.com/joaoasantana/e-product-service/internal/v1/infra/store"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(router *gin.RouterGroup, dbConn *mongo.Database) {
	categoryRepo := store.NewMongoCategoryRepository(dbConn)
	productRepo := store.NewMongoProductRepository(dbConn)
	productService := product.NewService(categoryRepo, productRepo)

	h := newHandler(productService)

	v1 := router.Group("/products")
	{
		v1.POST("/", h.createProduct)

		v1.GET("/", h.findAllProducts)
		v1.GET("/:id", h.findProductByID)
	}
}
