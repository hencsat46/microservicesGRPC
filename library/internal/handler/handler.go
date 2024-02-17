package handler

import (
	"context"
	"log"
	"net"

	library "github.com/hencsat46/protos/gen/go/library"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	library.UnimplementedLibServer
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	Create(string) (int, error)
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

	log.Println(username)

	_, err := h.usecase.Create(username)

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	return &library.RegisterResponse{Error: "NIL"}, nil
}
