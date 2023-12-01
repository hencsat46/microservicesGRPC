package handler

import (
	"context"
	"log"
	"net"

	auth "github.com/hencsat46/protos/gen/go/auth"
	"google.golang.org/grpc"
)

type handler struct {
	auth.UnimplementedAuthServer
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
}

func New(u UsecaseInterfaces) *handler {
	return &handler{
		usecase: u,
	}
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

func (g *handler) register(gRPC *grpc.Server, u UsecaseInterfaces) {
	auth.RegisterAuthServer(gRPC, &handler{usecase: u})
}

func (h *handler) Create(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	log.Println(request.GetUsername(), request.GetPassword(), request.GetFirstName(), request.GetSecondName())
	return &auth.RegisterResponse{UserId: "1", Error: "nil"}, nil
}

func (h *handler) Read(ctx context.Context, request *auth.ReadRequest) (*auth.ReadResponse, error) {
	panic("oops")
}

func (h *handler) Update(ctx context.Context, request *auth.UpdateRequest) (*auth.UpdateResponse, error) {
	panic("oops")
}

func (h *handler) Delete(ctx context.Context, request *auth.DeleteRequest) (*auth.DeleteResponse, error) {
	panic("oops")
}
