package entities

import "time"

// TIRT => Trimmed Iranian Toman, Means 1+3 less zeros compared to IRR (Iranian Rial)
type TIRTCharge struct {
	ID        int64 `bun:",pk,autoincrement"`
	UserID    int64
	Amount    float64
	CreatedAt time.Time `bun:",default:current_timestamp"`
}

type CryptoChargeStatus int

const (
	CCStatusCreated = iota
	CCStatusApplied
	CCStatusFailed
)

type ChainType int

const (
	ChainTypeTON = iota
	ChainTypeTRC
)

type CryptoCharge struct {
	ID        int64 `bun:",pk,autoincrement"`
	UserID    int64
	TxID      string `bun:",unique"`
	Amount    float64
	Chain     ChainType
	Status    CryptoChargeStatus
	CreatedAt time.Time `bun:",default:current_timestamp"`
	AppliedAt time.Time
}
