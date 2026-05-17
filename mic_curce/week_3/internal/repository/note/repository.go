package note

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/model"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository"
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note/converter"
	modelRepo "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note/model"
	//	desc "github.com/vivaldi7/golang_code/mic_curce/week_3/pkg/note_v1"
)

const (
	tableName       = "note"
	idColumn        = "id"
	titleColumn     = "title"
	contentColumn   = "content"
	craetedAtColumn = "created_at"
	updateAtColumn  = "update_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.NoteRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.NoteInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn, contentColumn).
		Values(info.Title, info.Content).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var id int64

	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		log.Fatalf("failed to insert notes: %v", err)
	}
	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Note, error) {
	builder := sq.Select(idColumn, titleColumn, contentColumn, craetedAtColumn, updateAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		log.Fatalf("Failed to build query: %v", err)
	}

	var note modelRepo.Note
	note.Info = &modelRepo.NoteInfo{}

	err = r.db.QueryRow(ctx, query, args...).Scan(&note.ID, &note.Info.Title, &note.Info.Content, &note.CreatedAt, &note.UpdateAt)
	if err != nil {
		log.Fatalf("Failed to select note: %v", err)
	}

	return converter.ToNoteFromRepo(&note), nil
}
