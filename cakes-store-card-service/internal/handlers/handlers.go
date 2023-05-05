package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/Cakes-store-card-service/entities"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServiceCard interface {
	AddItemToCart(userId int, Cake entities.Cake) error
	GetCart(userId int) (map[string]interface{}, error)
	Auth() gin.HandlerFunc
	VerifyToken(tokenSigned string) (string, error)
	GetToken(context *gin.Context) (string, error)
}

type AuthHandlers struct {
	ServiceCard ServiceCard
}

func NewAuthHandlers(serviceCard ServiceCard) *AuthHandlers {
	return &AuthHandlers{ServiceCard: serviceCard}
}

func SetRequestHandlers(serviceCard ServiceCard) (*gin.Engine, error) {
	router := gin.New()
	handlers := NewAuthHandlers(serviceCard)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(CORSMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello message")
	})
	cart := router.Group("/cart")
	{
		cart.POST("/addCart", handlers.addToCart)
		cart.GET("/getCart", handlers.cart).Use(serviceCard.Auth())
	}
	return router, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
