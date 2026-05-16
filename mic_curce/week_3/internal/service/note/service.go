package note

import (
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository"
	def "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/service"
)

var _ def.NoteService = (*serv)(nil)

type serv struct {
	noteRepository repository.NoteRepository
}

func NewService(noteRepository repository.NoteRepository) *serv {
	return &serv{
		noteRepository: noteRepository,
	}
}
