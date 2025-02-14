package godb

import (
	"context"
	"github.com/mrkucher83/avito-shop/internal/models"
)

func (i *Instance) GetInventoryById(ctx context.Context, employeeID int) ([]models.Purchase, error) {
	var inventory []models.Purchase

	query := `
		SELECT m.name AS type, p.quantity
		FROM purchases p
		JOIN merch m ON p.merch_id = m.id
		WHERE p.employee_id = $1
	`

	rows, err := i.Db.Query(ctx, query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var purchase models.Purchase
		if err = rows.Scan(&purchase.Type, &purchase.Quantity); err != nil {
			return nil, err
		}
		inventory = append(inventory, purchase)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return inventory, nil
}

func (i *Instance) GetReceivedCoins(ctx context.Context, employeeID int) ([]models.ReceiveCoin, error) {
	var received []models.ReceiveCoin

	query := `
		SELECT e.username, t.amount 
		FROM transactions t 
		JOIN employee e ON t.sender_id = e.id
		WHERE t.receiver_id = $1
	`

	rows, err := i.Db.Query(ctx, query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.ReceiveCoin
		if err = rows.Scan(&transaction.FromUser, &transaction.Amount); err != nil {
			return nil, err
		}
		received = append(received, transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return received, nil
}

func (i *Instance) GetSentCoins(ctx context.Context, employeeID int) ([]models.SendCoin, error) {
	var sent []models.SendCoin

	query := `
		SELECT e.username, t.amount 
		FROM transactions t
		JOIN employee e ON t.receiver_id = e.id
		WHERE t.sender_id = $1
	`

	rows, err := i.Db.Query(ctx, query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.SendCoin
		if err = rows.Scan(&transaction.ToUser, &transaction.Amount); err != nil {
			return nil, err
		}
		sent = append(sent, transaction)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sent, nil
}
