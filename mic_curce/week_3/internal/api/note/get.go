package note

import (
	"context"
	"log"

	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/converter"
	desc "github.com/vivaldi7/golang_code/mic_curce/week_3/pkg/note_v1"
)

func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	noteObj, err := s.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %v, title: %v, content: %v, created_at: %v, update_at: %v", noteObj.ID, noteObj.Info.Title, noteObj.Info.Content, noteObj.CreatedAt, noteObj.UpdateAt)

	return &desc.GetResponse{
		Note: converter.ToNoteFromService(noteObj),
	}, nil
}
