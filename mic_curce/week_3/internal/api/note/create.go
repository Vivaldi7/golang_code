package note

import (
	"context"
	"log"

	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/converter"
	desc "github.com/vivaldi7/golang_code/mic_curce/week_3/pkg/note_v1"
)

func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}
	log.Printf("Inserted note with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
