package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-product-service/internal/application/model"
	"github.com/joaoasantana/e-product-service/internal/application/service"
	"github.com/joaoasantana/e-product-service/pkg/util/response"
	"net/http"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	var requestBody model.ProductInput

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Failure{
			Status: response.Status{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: "Invalid request body",
		})
		return
	}

	if err := h.productService.Create(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.Failure{
			Status: response.Status{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response.Success{
		Status: response.Status{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		},
		Data: "Product created",
	})
}

func (h *ProductHandler) FindAll(ctx *gin.Context) {
	products, err := h.productService.FindAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.Failure{
			Status: response.Status{
				Code:    http.StatusNotFound,
				Message: http.StatusText(http.StatusNotFound),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Success{
		Status: response.Status{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: products,
	})
}

func (h *ProductHandler) FindByID(ctx *gin.Context) {
	product, err := h.productService.FindByID(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, response.Failure{
			Status: response.Status{
				Code:    http.StatusNotFound,
				Message: http.StatusText(http.StatusNotFound),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Success{
		Status: response.Status{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: product,
	})
}
