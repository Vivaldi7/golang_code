package repository

import (
	"context"

	desc "github.com/vivaldi7/golang_code/mic_curce/week_3/pkg/note_v1"
)

type NoteRepositoty interface {
	Create(ctx context.Context, info *desc.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*desc.Note, error)
}
