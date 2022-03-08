package main

import (
	"context"
	"database/sql"
	"log"
	"mify_api_radius/ent"
	"mify_api_radius/ent/proto/entpb"
	"net"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
)

func main() {
	conString := "postgres://radius:radius@localhost:5432/radius?sslmode=disable"
	db, err := sql.Open("pgx", conString)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	// Your code. For example:
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server := grpc.NewServer()

	svcNas := entpb.NewNasService(client)
	entpb.RegisterNasServiceServer(server, svcNas)

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Failed listening: %s", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatal("Server ended: %s", err)
	}
}
