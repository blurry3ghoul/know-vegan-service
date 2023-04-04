package controller

import (
	"math"
	"net/http"
	"strconv"

	"github.com/blurry3ghoul/know-vegan-service/src/api/domain"
	"github.com/blurry3ghoul/know-vegan-service/src/api/service"
	"github.com/blurry3ghoul/know-vegan-service/src/api/utils"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService service.ICategoryService
}

func NewCategoryController(categoryService service.ICategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var (
		category *domain.Category
	)

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	cc.categoryService.CreateCategory(category)

	c.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "",
		Data:    category,
	})
}

func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	// Get the page and limit parameters from the query string
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid 'page' field",
			Data:    nil,
		})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid 'limit' field",
			Data:    nil,
		})
		return
	}

	// Calculate the offset based on the page and limit parameters
	offset := (page - 1) * limit

	categories, total := cc.categoryService.GetAllCategories(offset, limit)

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	if categories == nil {
		c.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "Categories not found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.ResponsePagination{
		Success: true,
		Metadata: utils.Metadata{
			Page:       page,
			Limit:      limit,
			Total:      int(total),
			TotalPages: totalPages,
		},
		Data: categories,
	})
}
