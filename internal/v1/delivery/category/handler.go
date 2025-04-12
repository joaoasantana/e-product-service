package category

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-product-service/internal/v1/domain/app/category"
	"github.com/joaoasantana/e-product-service/pkg/util/response"
	"net/http"
)

type handler struct {
	categoryService *category.Service
}

func newHandler(categoryService *category.Service) *handler {
	return &handler{
		categoryService: categoryService,
	}
}

func (h *handler) createCategory(ctx *gin.Context) {
	var requestBody category.Input

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

	id, err := h.categoryService.Create(&requestBody)
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

func (h *handler) findAllCategories(ctx *gin.Context) {
	outputs, err := h.categoryService.FindAll()
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

func (h *handler) findCategoryByID(ctx *gin.Context) {
	output, err := h.categoryService.FindByID(ctx.Param("id"))
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
