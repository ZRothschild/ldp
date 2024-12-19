package main

import (
	"context"
	"github.com/ZRothschild/ldp/infrastr/conf"
	"github.com/ZRothschild/ldp/infrastr/mysql"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/ZRothschild/ldp/gen/user"            // Update
	userSrv "github.com/ZRothschild/ldp/server/user" // Update
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

/*
这个是支持 opts 中间键
1. grpc server 创建监听
2. http 请求创建
3. grpc client 创建
4. 当请求http http转发到grpc client 由他发送请求到 grpc server

or

当然也可以
1. grpc server 创建监听
2. http 请求创建
3. 当请求http http转发到 grpc server
*/

func run() error {
	var (
		err  error
		l    net.Listener
		addr = ":" + strconv.Itoa(conf.Conf.GrpcPort)
	)

	if l, err = net.Listen("tcp", addr); err != nil {
		log.Printf("Failed to listen: %v", err)
		return err
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	user.RegisterUserServiceServer(s, userSrv.NewUserServer())

	go func() {
		log.Fatalln(s.Serve(l))
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	// 这里是直接监听 Server
	if err = user.RegisterUserServiceHandlerServer(ctx, mux, userSrv.NewUserServer()); err != nil {
		log.Printf("Failed RegisterUserServiceHandlerServer: %v", err)
		return err
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//opts := []grpc.DialOption{grpc.WithInsecure()}
	// 这里是先通过 cli 在请求Server
	if err = user.RegisterUserServiceHandlerFromEndpoint(ctx, mux, addr, opts); err != nil {
		log.Printf("Failed RegisterUserServiceHandlerFromEndpoint: %v", err)
		return err
	}

	addr = ":" + strconv.Itoa(conf.Conf.HttpPort)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(addr, mux)
}

func main() {
	mysql.NewDb(conf.Conf)
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
