package google

import (
	"fmt"

	"github.com/a8uhnf/map-test/api"
)

func init() {
	fmt.Println("Hello google.com")
}

type SearchPlaces struct{}

func (s SearchPlaces) Search(*api.SearchPlacesRequest) (*api.SearchPlacesResponse, error) {
	ret := &api.SearchPlacesResponse{}
	
	return ret, nil
}
