package godb

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mrkucher83/avito-shop/internal/models"
)

func (i *Instance) GetMerchByName(ctx context.Context, itemName string) (*models.Merch, error) {
	query := `SELECT id, name, price FROM merch WHERE name = $1`
	var item models.Merch
	err := i.Db.QueryRow(ctx, query, itemName).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (i *Instance) GetCoinsById(ctx context.Context, employeeID int) (int, error) {
	query := `SELECT coins FROM employee WHERE id = $1`
	var coins int
	err := i.Db.QueryRow(ctx, query, employeeID).Scan(&coins)
	if err != nil {
		return 0, err
	}
	return coins, nil
}

func (i *Instance) BeginTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := i.Db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (i *Instance) UpdateBalance(ctx context.Context, tx pgx.Tx, price int, employeeID int) error {
	query := `UPDATE employee SET coins = coins - $1 WHERE id = $2`
	_, err := tx.Exec(ctx, query, price, employeeID)
	if err != nil {
		return err
	}
	return nil
}

func (i *Instance) RecordPurchase(ctx context.Context, tx pgx.Tx, employeeID, merchID, quantity int) error {
	query := `INSERT INTO purchases (employee_id, merch_id, quantity) VALUES ($1, $2, $3)`
	_, err := tx.Exec(ctx, query, employeeID, merchID, quantity)
	if err != nil {
		return err
	}
	return nil
}
