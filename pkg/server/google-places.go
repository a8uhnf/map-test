package server

import (
	"fmt"

	"github.com/a8uhnf/map-test/pkg/google"
)

func init() {
	fmt.Println("Hello from init func of searchPlaces pkg")
	MapToVendor["google"] = &google.SearchPlaces{}
}
