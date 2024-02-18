package handler

import (
	"context"
	"log"
	"net"

	"github.com/hencsat46/protos/gen/go/auth"
	lib "github.com/hencsat46/protos/gen/go/library"
	"github.com/hencsat46/protos/gen/go/manager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type handler struct {
	manager.UnimplementedManagerServer
}

type UsecaseInterfaces interface {
}

func New() *handler {
	return &handler{}
}

func (h *handler) Run(port string) error {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Println(err)
		return err
	}

	var opts []grpc.ServerOption

	server := grpc.NewServer(opts...)
	h.register(server)
	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (h *handler) register(gRPC *grpc.Server) {
	manager.RegisterManagerServer(gRPC, &handler{})
}

func (h *handler) CreateAccount(ctx context.Context, request *manager.RegisterRequest) (*manager.RegisterResponse, error) {
	// user := models.User{Username: request.Username, FirstName: request.FirstName, SecondName: request.SecondName, Password: request.Password}

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":3000", opts...)
	if err != nil {
		return &manager.RegisterResponse{UserId: -1, Status: "Connection to first microservice error"}, status.Error(codes.Unavailable, "cannot connect to microservice")
	}

	client := auth.NewAuthClient(conn)

	authRequest, err := client.Create(context.Background(), &auth.RegisterRequest{Username: request.GetUsername(), FirstName: request.GetFirstName(), SecondName: request.GetSecondName(), Password: request.GetPassword()})

	if err != nil {
		log.Println(err)
		return &manager.RegisterResponse{UserId: -1, Status: "Some shit"}, nil
	}

	//log.Println(authRequest.GetUserId())

	return &manager.RegisterResponse{Status: "Sign up ok", UserId: authRequest.GetUserId()}, nil
}

func (h *handler) CreateLibraryAccount(ctx context.Context, request *manager.RegisterLibRequest) (*manager.RegisterLibResponse, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":3500", opts...)
	if err != nil {
		return &manager.RegisterLibResponse{Status: "Cannot connect to microservice"}, status.Error(codes.Unavailable, "Cannot connect to microservice")
	}

	client := lib.NewLibClient(conn)

	libraryRequest, err := client.Add(context.Background(), &lib.RegisterRequest{Username: request.GetUsername()})
	if err != nil {
		log.Println(err)
		return &manager.RegisterLibResponse{Status: "Server error"}, err
	}

	return &manager.RegisterLibResponse{Status: libraryRequest.GetStatus()}, nil
}
