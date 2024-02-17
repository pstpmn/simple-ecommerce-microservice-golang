package pkg

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	// server struct {
	// 	pb.CustomerServiceServer
	// }

	IGrpc interface {
		Server(key, host string) (*grpc.Server, net.Listener)
		Client(host string) (*grpc.ClientConn, error)
	}

	g struct {
	}

	grpcAuth struct {
		secretKey string
	}
)

// Client implements IGrpc.
func (*g) Client(host string) (*grpc.ClientConn, error) {
	opts := make([]grpc.DialOption, 0)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(
		host,
		opts...,
	)
	return conn, err
}

// Server implements IGrpc.
func (*g) Server(key string, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)
	grpcAuth := &grpcAuth{
		secretKey: key,
	}
	opts = append(opts, grpc.UnaryInterceptor(grpcAuth.unaryAuthorization))
	grpcServer := grpc.NewServer(opts...)
	// pb.RegisterCustomerServiceServer(grpcServer, server{})
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}
	return grpcServer, lis
}

func NewGrpc() IGrpc {
	return &g{}
}

func (g *grpcAuth) unaryAuthorization(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// md, ok := metadata.FromIncomingContext(ctx)
	// fmt.Println(md, ok)
	// authHeader, ok := md["auth"]
	// fmt.Println(authHeader, ok)
	// if !ok {
	// 	log.Printf("Error: Metadata not found")
	// 	return nil, errors.New("error: metadata not found")
	// }

	// authHeader, ok := md["auth"]
	// if !ok {
	// 	log.Printf("Error: Metadata not found")
	// 	return nil, errors.New("error: metadata not found")
	// }

	// if len(authHeader) == 0 {
	// 	log.Printf("Error: Metadata not found")
	// 	return nil, errors.New("error: metadata not found")
	// }
	return handler(ctx, req)
}
