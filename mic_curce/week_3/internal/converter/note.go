package converter

import (
	"github.com/vivaldi7/golang_code/mic_curce/week_3/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/vivaldi7/golang_code/mic_curce/week_3/pkg/note_v1"
)

func ToNoteFromService(note *model.Note) *desc.Note {
	var updateAt *timestamppb.Timestamp
	if note.UpdateAt.Valid {
		updateAt = timestamppb.New(note.UpdateAt.Time)
	}

	return &desc.Note{
		Id:        note.ID,
		Info:      ToNoteInfoFromService(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdateAt:  updateAt,
	}
}

func ToNoteInfoFromService(info model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}

func ToNoteInfoFromDesc(info *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
