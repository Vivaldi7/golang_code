package converter

import (
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/model"
	modelRepo "github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note/model"
)

func ToNoteFromRepo(note *modelRepo.Note) *model.Note {

	return &model.Note{
		ID:        note.ID,
		Info:      ToNoteInfoFromRepo(note.Info),
		CreatedAt: note.CreatedAt,
		UpdateAt:  note.UpdateAt,
	}

}

func ToNoteInfoFromRepo(info *modelRepo.NoteInfo) model.NoteInfo {
	return model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}

}
