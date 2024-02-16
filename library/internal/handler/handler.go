package handler

import (
	library "github.com/hencsat46/protos/gen/go/library"
)

type handler struct {
	library.UnimplementedLibraryServiceServer
	usecase UsecaseInterfaces
}

type UsecaseInterfaces interface {
	
}