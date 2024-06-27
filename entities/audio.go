package entities

import "time"

type Audio struct {
	ID            int64 `bun:",pk,autoincrement"`
	UserID        int64
	FileName      string
	FileUniqueID  string
	MessageID     int64
	Transcription string
	Translation   string
	CreatedAt     time.Time `bun:",default:current_timestamp"`
}
