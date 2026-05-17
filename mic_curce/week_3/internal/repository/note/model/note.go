package modelRepo

import (
	"database/sql"
	"time"
)

type Note struct {
	ID        int64        `db:"id"`
	Info      *NoteInfo    `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdateAt  sql.NullTime `db:"update_at"`
}

type NoteInfo struct {
	Title   string `db:"title"`
	Content string `db:"content"`
}
