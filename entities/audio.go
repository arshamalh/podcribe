package entities

import "time"

type Audio struct {
	ID            int64 `bun:",pk,autoincrement"`
	UserID        int64
	FileName      string
	FileUniqueID  string
	MessageID     int
	Duration      int
	Transcription string // TODO: Transcription and Translation are long texts, they are might not be suitable to be stored in this column.
	Translation   string
	CreatedAt     time.Time `bun:",default:current_timestamp"`
}

func NewAudio(userID int64, filename string, fileUniqueID string, messageID int) *Audio {
	return &Audio{
		UserID:       userID,
		FileName:     filename,
		FileUniqueID: fileUniqueID,
		MessageID:    messageID,
	}
}

func (a *Audio) IssueInvoice(factor float64) *Invoice {
	return &Invoice{
		AudioID: a.ID,
		UserID:  a.UserID,
		Amount:  float64(a.Duration) * factor,
	}
}
