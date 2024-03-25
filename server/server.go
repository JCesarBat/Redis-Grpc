package server

import (
	"Redis/pb"
	database "Redis/store/sqlc"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	arg := database.CreatePersonaParams{
		Nombre:    request.GetNombre(),
		Ocupacion: request.GetOcupacion(),
		Edad:      request.GetEdad(),
	}
	fmt.Println(request.GetNombre(), "wewweewe")
	person, err := server.store.CreatePersona(context.Background(), arg)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.Aborted, "error ocurred in database ")
		}
		return nil, status.Errorf(codes.Internal, "cannot create the person :", err)
	}

	response := &pb.CreatePersonResponse{
		Persona: &pb.Persona{
			Nombre:    person.Nombre,
			Edad:      uint32(person.Edad),
			Ocupacion: person.Ocupacion,
		},
	}
	return response, nil
}
