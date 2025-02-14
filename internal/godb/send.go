package godb

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func (i *Instance) InsertTransaction(ctx context.Context, tx pgx.Tx, senderID, receiverID, amount int) error {
	query := `INSERT INTO transactions (sender_id, receiver_id, amount) VALUES ($1, $2, $3)`
	_, err := tx.Exec(ctx, query, senderID, receiverID, amount)
	if err != nil {
		return err
	}
	return nil
}
