package routes

import (
	"github.com/blurry3ghoul/know-vegan-service/src/api/controller"
	"github.com/blurry3ghoul/know-vegan-service/src/api/service"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, productService service.IProductService) {
	c := controller.NewProductController(productService)

	r.POST("/v1/product", c.CreateProduct)
}
