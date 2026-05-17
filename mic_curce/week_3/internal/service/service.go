package service

import (
	"context"

	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/model"
)

type NoteService interface {
	Create(ctx context.Context, info *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
}
