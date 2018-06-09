package api

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct{}

var count int32

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *SearchPlacesRequest) (*SearchPlacesResponse, error) {
	ret := &SearchPlacesResponse{}
	env := os.Getenv("HELLO_WORLD")

	fmt.Println("--------------- ENV: ", env)
	// n, err := strconv.Atoi(env)
	fmt.Println("--------------- COUNT: ", count)
	count++

	log.Printf("Receive message %s", in.Descriptor)
	return ret, nil
}
