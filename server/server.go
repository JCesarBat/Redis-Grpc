package server

import (
	"Redis/pb"
	database "Redis/store/sqlc"
	"context"
)

type Server struct {
	pb.UnimplementedServerRpcServer
	store *database.Store
}

func NewServer(store *database.Store) *Server {

	return &Server{
		store: store,
	}
}

func (server Server) CreatePerson(ctx context.Context, request *pb.CreatePersonRequest) (*pb.CreatePersonResponse, error) {

	return nil, nil
}
