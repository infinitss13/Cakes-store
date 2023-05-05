package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitss13/Cakes-store-card-service/entities"
	"github.com/sirupsen/logrus"
)

func (handlers AuthHandlers) addToCart(context *gin.Context) {
	userId, exists := context.Get("userID")
	if !exists {
		logrus.Error(errors.New("user ID is not found "))
		context.AbortWithStatusJSON(http.StatusBadRequest, errors.New("user ID is not found "))
		return
	}
	cake := entities.NewCake()
	err := context.ShouldBindJSON(&cake)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logrus.Error(err)
		return
	}
	err = handlers.ServiceCard.AddItemToCart(userId.(int), cake)
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
	userId, exists := context.Get("userID")
	if !exists {
		logrus.Error(errors.New("user ID is not found "))
		context.AbortWithStatusJSON(http.StatusBadRequest, errors.New("user ID is not found "))
		return
	}
	//here must be implementation of getting user id from the header.

	cakes, err := handlers.ServiceCard.GetCart(userId.(int))
	if err != nil {
		logrus.Error(err)
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"userCart": cakes,
	})
}
