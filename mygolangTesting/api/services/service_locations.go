package services

import (
	"fmt"

	"github.com/selva2474/golang_testing/mygolangTesting/api/domain/locations/providerlocations"
	"github.com/selva2474/golang_testing/mygolangTesting/api/utils/errors"
	"github.com/selva2474/golang_testing/src/api/domain/locations"
)

// initialized = false set to true when the package is called
type locationsService struct{}

type locationServiceInterface interface {
	GetCountry(countryID string) (*locations.Country, *errors.APIerror)
}

// LocationsService is the locationServiceInterface
var LocationsService locationServiceInterface

func init() {
	fmt.Println("Init Service")
	LocationsService = &locationsService{}
}

// GetCountry takes a countryID
// and returns a pointer to locations.Country and a pointer to an APIerror
func (s *locationsService) GetCountry(countryID string) (*locations.Country, *errors.APIerror) {
	fmt.Println("Inside Service")
	return providerlocations.GetCountry(countryID)

}
