package handlers

import (
	"context"
	"podcribe/config"
	"podcribe/entities"
	"podcribe/telegram/keyboards"
	"podcribe/telegram/msgs"
)

func (h *handler) Credit(ctx SharedContext) error {
	user, err := h.db.GetUserByChatID(context.TODO(), ctx.ID())
	if err != nil {
		return ctx.Send("unexpected problem happened")
	}

	ctx.SetScene(entities.SceneCredit)

	// TODO:
	// Change scene and add a handler for getting user transactionID,
	// then lock the transaction ID and another cronjob will watch the blockchain for any incoming transaction to it

	return ctx.Send(
		msgs.FmtCredit(user.Balance),
		keyboards.Credit(),
	)
}

// TRON Transaction IDs are 64 characters length
// TON Transaction IDs are 44 characters length
func (h *handler) CreditTxTextHandler(ctx SharedContext) error {
	txID := ctx.Text()
	if len(txID) == 44 {
		h.ton.CheckTransaction(txID)
		// First check if the transaction doesn't exist
		// If the transaction was valid
	} else if len(txID) == 64 {
		h.tron.CheckTransaction(txID)
	} else {
		return ctx.Send("invalid transaction ID, contact admin: " + config.Get().AdminUsername)
	}

	return nil
}
