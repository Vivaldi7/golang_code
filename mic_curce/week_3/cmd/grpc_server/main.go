package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/config"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note"
	desc "github.com/vivaldi7/golang_code/mic_curce/week_3/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	desc.UnimplementedNoteV1Server
	noteRepository repository.NoteRepository
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.noteRepository.Create(ctx, req.GetInfo())
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted note with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	noteObj, err := s.noteRepository.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %v, title: %v, body: %v, created_at: %v, update_at: %v", noteObj.Id, noteObj.Info.Title, noteObj.Info.Content, noteObj.CreatedAt, noteObj.UpdateAt)

	return &desc.GetResponse{
		Note: noteObj,
	}, nil
}

func main() {

	ctx := context.Background()

	//Считываем переенные окружения
	err := config.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.GRPCAddress())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to conect to database: %v", err)
	}

	defer pool.Close()

	noteRepo := note.NewRepository(pool)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{noteRepository: noteRepo})

	log.Printf("Server listening at: %v", grpcConfig.GRPCAddress())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
