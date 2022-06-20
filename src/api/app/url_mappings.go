package app

import (
	"github.com/selva2474/golang_testing/src/api/controllers"
)

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
