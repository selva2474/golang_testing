package app

import "github.com/selva2474/golang_testing/src/api/controllers"

// import "github.com/golang-testing/mygolangTesting/api/controllers"

func mapURLs() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
