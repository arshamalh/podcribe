package sqlite

import (
	"context"
	"database/sql"
	"podcribe/entities"
	"podcribe/log"

	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// Make a new invoice and decreases user balance
func (s *Sqlite) AddInvoice(ctx context.Context, invoice *entities.Invoice) error {
	err := s.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(invoice).Exec(ctx); err != nil {
			return err
		}

		_, err := tx.NewUpdate().
			Model(&entities.User{}).
			Where("id = ?", invoice.UserID).
			Set("balance = balance - ?", invoice.Amount).
			Exec(ctx)

		return err
	})

	if err != nil {
		return err
	}

	log.Gl.Info("new invoice created:", zap.Int64("userID", invoice.UserID), zap.Int64("invoiceID", invoice.ID))
	return nil
}
