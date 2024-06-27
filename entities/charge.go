package entities

import "time"

type ChargeType int

const (
	ChargeTypeTON = iota
	ChargeTypeTIRT
)

type Charge struct {
	ID        int64 `bun:",pk,autoincrement"`
	UserID    int64
	Type      ChargeType
	CreatedAt time.Time `bun:",default:current_timestamp"`
}

type TONCharge struct {
	ID       int64 `bun:",pk,autoincrement"`
	ChargeID int64
	Amount   float64
}

// TIRT => Trimmed Iranian Toman, Means 1+3 less zeros compared to IRR (Iranian Rial)
type TIRTCharge struct {
	ID       int64 `bun:",pk,autoincrement"`
	ChargeID int64
	Amount   float64
}
