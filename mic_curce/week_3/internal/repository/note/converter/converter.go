package converter

import (
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/repository/note/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/vivaldi7/golang_code/mic_curce/week_3/pkg/note_v1"
)

func ToNoteFromRepo(note *model.Note) *desc.Note {
	var updateAt *timestamppb.Timestamp
	if note.UpdateAt.Valid {
		updateAt = timestamppb.New(note.UpdateAt.Time)
	}

	return &desc.Note{
		Id:        note.ID,
		Info:      ToNoteinfoFromRepo(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdateAt:  updateAt,
	}

}

func ToNoteinfoFromRepo(info *model.Info) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}

}
