package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
	desc "github.com/vivaldi7/golang_code/mic_curce/week_1/grpc/pkg/note_v1"
	"github.com/vivaldi7/golang_code/mic_curce/week_2/config/internal/config"
	"github.com/vivaldi7/golang_code/mic_curce/week_2/config/internal/config/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedNoteV1Server
	pool *pgxpool.Pool
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	//Делаем запрос на вставку записи в таблице note
	builderInsert := sq.Insert("note").
		PlaceholderFormat(sq.Dollar).
		Columns("title", "body").
		Values(gofakeit.City(), gofakeit.Address().Street).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var noteID int64

	err = s.pool.QueryRow(ctx, query, args...).Scan(&noteID)
	if err != nil {
		log.Fatalf("failed to insert note: %v", err)
	}

	log.Printf("Inserted note with id: %d", noteID)

	return &desc.CreateResponse{Id: noteID}, nil

}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	//Делаем запрос на запись по ID
	builderSelectOne := sq.Select("id", "title", "body", "created_at", "update_at").
		From("note").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}
	var id int64
	var title, body string
	var createdAt time.Time
	var updateAt sql.NullTime

	err = s.pool.QueryRow(ctx, query, args...).Scan(&id, &title, &body, &createdAt, &updateAt)
	if err != nil {
		log.Fatalf("failed to scan note: %v", err)
	}

	var updateAtTime *timestamppb.Timestamp
	if updateAt.Valid {
		updateAtTime = timestamppb.New(updateAt.Time)
	}

	log.Printf("id: %v, title: %v, body: %v, created_at: %v, update_at: %v", id, title, body, createdAt, updateAt)

	return &desc.GetResponse{
		Note: &desc.Note{
			Id:        id,
			Info:      &desc.NoteInfo{Title: title, Content: body},
			CreatedAt: timestamppb.New(createdAt),
			UpdateAt:  updateAtTime,
		},
	}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	//Считываем переенные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to conect to database: %v", err)
	}

	defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{pool: pool})

	log.Printf("Server listening at: %v", grpcConfig.Address())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
