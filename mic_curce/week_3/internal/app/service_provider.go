package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/api/note"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/closer"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/config"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository"
	noteRepositore "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/service"

	//	note "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/api/note"
	//	noteRepositore "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note"
	//	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note"
	//	noteRepositore "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note"
	//	noteService "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/service/note"
	//	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/service/note"
	noteService "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/service/note"
	//	noteService "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/service/note"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	pgPool     *pgxpool.Pool

	noteRepository repository.NoteRepository
	noteService    service.NoteService
	noteImpl       *note.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) PGPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}
	return s.pgPool
}

func (s *serviceProvider) NoteRepositore(ctx contex.Context) repositore.NoteRepositore {
	if s.noteRepositore == nil {
		s.noteRepositore = noteRepositore.NewRepositore(s.PGPool(ctx))
	}
}

func (s *serviceProvider) NoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = noteService.NewService(s.NoteRepositore(ctx))
	}
}

func (s *serviceProvider) noteImpl(ctx context.Context) *note.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = note.NewImplementation(s.NoteService(ctx))
	}
}
