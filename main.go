package main

import (
	"Redis/pb"
	"Redis/server"
	database "Redis/store/sqlc"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

type Person struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Ocupation string `json:"ocupation"`
}

func main() {
	conn, err := sql.Open("postgres", "postgresql://Jc:secret@localhost:5432/redisdb?sslmode=disable")

	if err != nil {
		fmt.Print("cannot connect to database: ", err)
		return
	}
	store := database.NewStore(conn)
	go StartGateWey(store, context.Background())
	StartGrpcServer(store)

	/*
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		ping, err := client.Ping(context.Background()).Result()

		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(ping)

		jsonString, err := json.Marshal(Person{
			Name:      "Eliot",
			Age:       22,
			Ocupation: "Programmer"})

		err = client.Set(context.Background(), "Persona", jsonString, 0).Err()

		if err != nil {
			fmt.Printf("failed to set value in the redis instance %s", err.Error())
			return
		}

		val, err := client.Get(context.Background(), "Persona").Result()

		if err != nil {
			fmt.Printf("failed to get value from redis %s", err.Error())
			return
		}

		fmt.Println(val)
	*/
}

func StartGrpcServer(store *database.Store) {
	server := server.NewServer(store)

	grpcServer := grpc.NewServer()

	pb.RegisterServerRpcServer(grpcServer, server)
	// Register gRPC services here if needed
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("gRPC Server is listening on", "0.0.0.0:9000")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func StartGateWey(store *database.Store, ctx context.Context) {
	server := server.NewServer(store)

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	err := pb.RegisterServerRpcHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("error ocurred in gatewey:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatal("error ocurred listening the gatewey :", err)
	}

	fmt.Println("http server gatewey start in port  :8000")
	err = http.Serve(listener, mux)

	if err != nil {
		log.Fatal("cannot start http gatewey")
	}

}
