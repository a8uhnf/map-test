package server

import (
	"fmt"
	"log"
	"os"

	"github.com/a8uhnf/map-test/api"
	"golang.org/x/net/context"
)

var MapToVendor map[string]interface{}

// Server represents the gRPC server
type Server struct{}

var count int32

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *api.SearchPlacesRequest) (*api.SearchPlacesResponse, error) {
	ret := &api.SearchPlacesResponse{}
	env := os.Getenv("HELLO_WORLD")

	fmt.Println("--------------- ENV: ", env)
	// n, err := strconv.Atoi(env)
	fmt.Println("--------------- COUNT: ", count)
	count++

	log.Printf("Receive message %s", in.Descriptor)
	return ret, nil
}
