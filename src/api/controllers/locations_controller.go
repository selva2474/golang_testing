package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selva2474/golang_testing/src/api/services"
)

func GetCountry(c *gin.Context) {
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
