package server

import (
	"fmt"
	"log"
	"os"

	"github.com/a8uhnf/map-test/api"
	"github.com/kr/pretty"
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
	log.Println("--------")
	pretty.Println(in)
	ret := &api.SearchPlacesResponse{}
	env := os.Getenv("HELLO_WORLD")

	fmt.Println("--------------- ENV: ", env)
	// n, err := strconv.Atoi(env)
	fmt.Println("--------------- COUNT: ", count)
	count++

	log.Printf("Receive message %s", in.Descriptor)

	vendor := in.Vendor
	intf := MapToVendor[vendor]
	log.Println("This is the vendor----", vendor)

	resp, err := intf.Search(in)
	if err != nil {
		return nil, err
	}
	pretty.Println(resp)

	return ret, nil
}
