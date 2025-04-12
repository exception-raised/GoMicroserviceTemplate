package server

import (
	"github.com/exception-raised/microservices-template/src/service"
	"github.com/exception-raised/microservices-template/src/storage"
)

type Option func(*Server)

func WithService(service *service.Service) Option {
	return func(server *Server) {
		server.service = service
	}
}

func WithStorage(storage storage.Client) Option {
	return func(server *Server) {
		server.storage = storage
	}
}
