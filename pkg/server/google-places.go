package server

import (
	"log"

	"github.com/a8uhnf/map-test/config"
	"github.com/a8uhnf/map-test/pkg/google"
)

func init() {
	g := &google.SearchPlaces{}
	cfg, err := config.GetConfig()
	g.Config = *cfg
	if err != nil {
		log.Fatalln(err)
	}
	MapToVendor["google"] = g
}
