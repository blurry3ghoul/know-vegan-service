package routes

import (
	"github.com/blurry3ghoul/know-vegan-service/src/api/controller"
	"github.com/blurry3ghoul/know-vegan-service/src/api/service"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine, categoryService service.ICategoryService) {
	c := controller.NewCategoryController(categoryService)

	r.POST("/v1/category", c.CreateCategory)
	r.GET("/v1/category", c.GetAllCategories)
}
