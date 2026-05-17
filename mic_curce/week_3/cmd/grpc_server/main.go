package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/config"

	//	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/converter"
	noteRepositore "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note"
	//	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/service"
	noteApi "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/api/note"
	noteService "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/service/note"
	desc "github.com/vivaldi7/golang_code/mic_curce/week_3/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

	noteRepo := noteRepositore.NewRepository(pool)
	noteSrv := noteService.NewService(noteRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, noteApi.NewImplementation(noteSrv))

	log.Printf("Server listening at: %v", grpcConfig.GRPCAddress())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
