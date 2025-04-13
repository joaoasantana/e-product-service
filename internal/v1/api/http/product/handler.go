package product

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/app/product"
	"github.com/joaoasantana/e-product-service/pkg/util/response"
	"net/http"
)

type handler struct {
	productService *product.Service
}

func newHandler(productService *product.Service) *handler {
	return &handler{
		productService: productService,
	}
}

func (h *handler) createProduct(ctx *gin.Context) {
	var requestBody product.Input

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.Failure{
			Status: response.Status{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	id, err := h.productService.Create(&requestBody)
	if err != nil {
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
		Data: id,
	})
}

func (h *handler) findAllProducts(ctx *gin.Context) {
	outputs, err := h.productService.FindAll()
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
		Data: outputs,
	})
}

func (h *handler) findProductByID(ctx *gin.Context) {
	output, err := h.productService.FindByID(ctx.Param("id"))
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
		Data: output,
	})
}
