package entities

import "time"

type Invoice struct {
	ID        int64 `bun:",pk,autoincrement"`
	UserID    int64
	AudioID   int64
	Amount    float64
	CreatedAt time.Time `bun:",default:current_timestamp"`
}
