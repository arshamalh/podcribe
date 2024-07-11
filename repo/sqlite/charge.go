package sqlite

import (
	"context"
	"podcribe/entities"
)

func (s *Sqlite) GetLatestCharge(ctx context.Context) (*entities.CryptoCharge, error) {
	charge := new(entities.CryptoCharge)
	err := s.db.
		QueryRowContext(ctx, "SELECT id, user_id, tx_id, amount, chain, status, created_at, applied_at FROM crypto_charges ORDER BY created_at LIMIT 1").
		Scan(
			&charge.ID, &charge.UserID, &charge.TxID,
			&charge.Amount, &charge.Chain, &charge.Status,
			&charge.CreatedAt, &charge.AppliedAt,
		)
	return charge, err
}
