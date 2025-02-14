package godb

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mrkucher83/avito-shop/pkg/logger"
)

func (i *Instance) BeginTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := i.Db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func RollbackOnError(ctx context.Context, tx pgx.Tx) {
	if err := tx.Rollback(ctx); err != nil {
		logger.Error("Failed to rollback transaction: %v", err)
	}
}
