package godb

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func (i *Instance) ReduceBalance(ctx context.Context, tx pgx.Tx, amount, employeeID int) error {
	query := `UPDATE employee SET coins = coins - $1 WHERE id = $2`
	_, err := tx.Exec(ctx, query, amount, employeeID)
	if err != nil {
		return err
	}
	return nil
}

func (i *Instance) IncreaseBalance(ctx context.Context, tx pgx.Tx, amount, employeeID int) error {
	query := `UPDATE employee SET coins = coins + $1 WHERE id = $2`
	_, err := tx.Exec(ctx, query, amount, employeeID)
	if err != nil {
		return err
	}
	return nil
}
