package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/fatih/color"
	desc "github.com/vivaldi7/golang_code/mic_curce/grpc/pkg/note_v1"
)

const (
	address = "localhost:50051"
	noteID  = 12
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to conect to server: %v", err)
	}

	defer conn.Close()

	c := desc.NewNoteV1Client(conn)

	ctx, cansel := context.WithTimeout(context.Background(), time.Second)
	defer cansel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Fatalf("failed to get Note by id: %v", err)
	}

	log.Printf(color.RedString("Note info:\n"), color.GreenString("%+v", r.GetNote()))
}
