package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoasantana/e-product-service/internal/application/model"
	"github.com/joaoasantana/e-product-service/internal/application/service"
	"github.com/joaoasantana/e-product-service/pkg/util/response"
	"net/http"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) Create(ctx *gin.Context) {
	var requestBody model.CategoryInput

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

	if err := h.categoryService.Create(&requestBody); err != nil {
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
		Data: "Category created",
	})
}

func (h *CategoryHandler) FindAll(ctx *gin.Context) {
	categories, err := h.categoryService.FindAll()
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
		Data: categories,
	})
}

func (h *CategoryHandler) FindByID(ctx *gin.Context) {
	category, err := h.categoryService.FindByID(ctx.Param("id"))
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
		Data: category,
	})
}
