package handler

import (
	"context"
	"log"
	"net"

	"github.com/hencsat46/protos/gen/go/auth"
	library "github.com/hencsat46/protos/gen/go/library"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type handler struct {
	library.UnimplementedLibServer
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	Create(string) error
	Get(string) bool
	Delete(string) error
}

func NewHandler(u UsecaseInterfaces) *handler {
	return &handler{usecase: u}
}

func (h *handler) Run(port string) error {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	server := grpc.NewServer(opts...)

	h.register(server, h.usecase)

	if err := server.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (h *handler) register(gRPC *grpc.Server, u UsecaseInterfaces) {
	library.RegisterLibServer(gRPC, &handler{usecase: u})
}

func (h *handler) Add(ctx context.Context, request *library.RegisterRequest) (*library.RegisterResponse, error) {
	username := request.GetUsername()

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":3000", opts...)
	if err != nil {
		log.Println(err)
		return &library.RegisterResponse{Status: "Connection to first microservice error"}, status.Error(codes.Unavailable, "cannot connect to microservice")
	}

	client := auth.NewAuthClient(conn)

	feature, err := client.Read(context.Background(), &auth.ReadRequest{Username: username})

	if err != nil {
		log.Println(err)
		return &library.RegisterResponse{Status: "Not found"}, err
	}

	log.Println(feature)

	if feature.Username == username {
		if err = h.usecase.Create(username); err != nil {
			return nil, status.Error(codes.Internal, "Internal Server Error")
		}
	}

	return &library.RegisterResponse{Status: "Registration ok"}, err

}

func (h *handler) Get(ctx context.Context, request *library.GetRequest) (*library.GetResponse, error) {
	username := request.GetUsername()

	if h.usecase.Get(username) {
		return &library.GetResponse{Status: "You are signed up"}, nil
	}
	return &library.GetResponse{Status: "Not found"}, status.Error(codes.NotFound, "Not Found")
}

func (h *handler) Delete(ctx context.Context, request *library.DeleteRequest) (*library.DeleteResponse, error) {
	username := request.GetUsername()

	if err := h.usecase.Delete(username); err != nil {
		return &library.DeleteResponse{Status: "Not Found"}, err
	}

	return &library.DeleteResponse{Status: "Delete Ok"}, nil
}
