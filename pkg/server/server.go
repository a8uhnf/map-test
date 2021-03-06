package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/a8uhnf/map-test/api"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// private type for Context keys
type contextKey int

const (
	clientIDKey contextKey = iota
)

func credMatcher(headerName string) (mdName string, ok bool) {
	if headerName == "Login" || headerName == "Password" {
		return headerName, true
	}
	return "", false
}

func startGRPCServer(address string) error {
	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	// create a server instance
	s := Server{}
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	api.RegisterSearchPlacesServer(grpcServer, &s)
	// start the server
	log.Printf("starting HTTP/2 gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}
	return nil
}

func startRESTServer(address, grpcAddress string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(credMatcher))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	// Register searchresults
	err := api.RegisterSearchPlacesHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return fmt.Errorf("could not register service Ping: %s", err)
	}
	log.Printf("starting HTTP/1.1 REST server on %s", address)
	http.ListenAndServe(address, mux)
	return nil
}

// StartServer start a gRPC server and waits for connection
func StartServer() {
	grpcAddress := fmt.Sprintf("%s:%d", "localhost", 7777)
	restAddress := fmt.Sprintf("%s:%d", "localhost", 7778)
	// fire the gRPC server in a goroutine
	go func() {
		err := startGRPCServer(grpcAddress)
		if err != nil {
			log.Fatalln("failed to start gRPC server: ", err)
		}
	}()
	// fire the REST server in a goroutine
	go func() {
		err := startRESTServer(restAddress, grpcAddress)
		if err != nil {
			log.Fatalln("failed to start REST server: ", err)
		}
	}()
	// infinite loop
	log.Printf("Entering infinite loop")
	select {}
}
