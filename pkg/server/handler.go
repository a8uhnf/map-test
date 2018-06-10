package server

import (
	"github.com/a8uhnf/map-test/api"
	"golang.org/x/net/context"
)

type SearchInterface interface {
	Search(*api.SearchPlacesRequest) (*api.SearchPlacesResponse, error)
}

var MapToVendor = make(map[string]SearchInterface)

// Server represents the gRPC server
type Server struct{}

var count int32

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *api.SearchPlacesRequest) (*api.SearchPlacesResponse, error) {
	vendor := in.Vendor
	intf := MapToVendor[vendor]
	resp, err := intf.Search(in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
