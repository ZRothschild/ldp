package main

import (
	"context"
	"flag"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"net/http"

	"github.com/ZRothschild/ldp/gen/user"            // Update
	userSrv "github.com/ZRothschild/ldp/server/user" // Update
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func run() error {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	user.RegisterUserServiceServer(s, userSrv.NewUserServer())
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
