package api

import (
	"github.com/blurry3ghoul/know-vegan-service/src/api/client/mysql"
	"github.com/blurry3ghoul/know-vegan-service/src/api/client/s3"
	"github.com/blurry3ghoul/know-vegan-service/src/api/config"
	"github.com/blurry3ghoul/know-vegan-service/src/api/middleware"
	"github.com/blurry3ghoul/know-vegan-service/src/api/persistance/repository"
	"github.com/blurry3ghoul/know-vegan-service/src/api/routes"
	"github.com/blurry3ghoul/know-vegan-service/src/api/service"
	"github.com/gin-gonic/gin"
)

func Execute() {
	gin.SetMode(config.MODE)

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	categoryRepository := repository.NewCategoryRepository(mysql.DB)
	productRepository := repository.NewProductRepository(mysql.DB)

	categoryService := service.NewCategoryService(categoryRepository)
	productService := service.NewProductService(productRepository, s3.GetSession)

	routes.ProductRoutes(r, productService)
	routes.CategoryRoutes(r, categoryService)

	err := r.Run(config.HTTPPORT)
	if err != nil {
		panic(err)
	}
}
