package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary catalog
// @Description handler for getting all the catalog from the database
// @ID catalog
// @Produce json
// @Success 200 {object} entities.Cake
// @Router /catalog [get]
func (h *CatalogHandlers) catalog(ctx *gin.Context) {
	limit := ctx.Query("limit")
	cakes, err := h.ServiceCatalog.GetCatalog(limit)
	if err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, cakes)
}

func (h *CatalogHandlers) addToCart(ctx *gin.Context) {
	id, exists := ctx.Get("userID")
	fmt.Println(id, exists)
}
