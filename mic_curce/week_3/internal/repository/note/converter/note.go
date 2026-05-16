package converter

import (
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note/model"

	desc "github.com/vivaldi7/golang_code/mic_curce/week_3/pkg/note_v1"
)

func ToNoteFromRepo(note *modelRepo.Note) *desc.Note {

	/*	var updateAt *timestamppb.Timestamp
		if note.UpdateAt.Valid {
			updateAt = timestamppb.New(note.UpdateAt.Time)
		}*/

	return &model.Note{
		Id:        note.ID,
		Info:      ToNoteinfoFromRepo(note.Info),
		CreatedAt: note.CreatedAt,
		UpdateAt:  note.UpdateAt,
	}

}

func ToNoteinfoFromRepo(info *modelRepo.Info) *model.NoteInfo {
	return &model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}

}
