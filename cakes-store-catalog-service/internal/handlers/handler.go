package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/Cakes-store-catalog-service/entities"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type CatalogHandlers struct {
	ServiceCatalog ServiceCatalog
}

type ServiceCatalog interface {
	GetCatalog() ([]entities.Cake, error)
}

func NewAuthHandlers(serviceAuth ServiceCatalog) *CatalogHandlers {
	return &CatalogHandlers{ServiceCatalog: serviceAuth}
}

func SetRequestHandlers(serviceAuth ServiceCatalog) (*gin.Engine, error) {
	router := gin.New()
	handlers := NewAuthHandlers(serviceAuth)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello message")
	})

	router.GET("/catalog", handlers.catalog)

	return router, nil
}
