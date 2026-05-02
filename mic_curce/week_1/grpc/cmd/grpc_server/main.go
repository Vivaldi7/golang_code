package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/brianvoe/gofakeit"

	desc "github.com/vivaldi7/golang_code/mic_curce/grpc/pkg/note_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedNoteV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Note id %d:", req.GetId())

	return &desc.GetResponse{
		Note: &desc.Note{
			Id: req.GetId(),
			Info: &desc.NoteInfo{
				Title:    gofakeit.BeerName(),
				Content:  gofakeit.IPv4Address(),
				Author:   gofakeit.Name(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdateAt:  timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})

	fmt.Printf("servr listening at: %v", grpcPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
