package server

import (
	"github.com/a8uhnf/map-test/pkg/google"
)

func init() {
	MapToVendor["google"] = &google.SearchPlaces{}
}
