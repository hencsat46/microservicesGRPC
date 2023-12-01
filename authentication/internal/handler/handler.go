package handler

import (
	"context"
	"log"
	"microservicesGRPC/authentication/internal/models"
	"net"

	auth "github.com/hencsat46/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	auth.UnimplementedAuthServer
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	Create(user *models.User) (int, error)
	Read(id int) (*models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
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

func (h *handler) register(gRPC *grpc.Server, u UsecaseInterfaces) {
	auth.RegisterAuthServer(gRPC, &handler{usecase: u})
}

func (h *handler) Create(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	user := models.User{Username: request.GetUsername(), Password: request.GetPassword(), FirstName: request.GetFirstName(), SecondName: request.GetSecondName()}

	id, err := h.usecase.Create(&user)
	if err != nil {
		log.Println("жопа")
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}
	id32 := int32(id)
	//log.Println((&auth.RegisterResponse{UserId: id32, Error: "nil"}).GetError())
	//log.Println(request.GetUsername(), request.GetPassword(), request.GetFirstName(), request.GetSecondName())
	return &auth.RegisterResponse{UserId: id32, Error: "nil"}, nil
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
