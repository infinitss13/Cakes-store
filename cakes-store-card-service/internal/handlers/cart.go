package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/Cakes-store-card-service/entities"
	"github.com/sirupsen/logrus"
)

func (handlers AuthHandlers) addToCart(context *gin.Context) {
	userCart := entities.NewUserCart()
	err := context.ShouldBindJSON(&userCart)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}
	err = handlers.ServiceCard.AddItemToCart(userCart)
	if err != nil {
		logrus.Error(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Item has been successfully added to cart",
	})
}

func (handlers AuthHandlers) cart(context *gin.Context) {
	userId := 123
	fmt.Println(userId)
	cakes, err := handlers.ServiceCard.GetCart(userId)
	if err != nil {
		logrus.Error(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"userCart": cakes,
	})
}
