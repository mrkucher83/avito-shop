package godb

import (
	"context"
	"github.com/mrkucher83/avito-shop/internal/models"
)

func (i *Instance) CreateEmployee(ctx context.Context, empl models.AuthRequest) error {
	query := `INSERT INTO employee (username, password, coins) VALUES ($1, $2, 1000);`

	_, err := i.Db.Exec(ctx, query, empl.Username, empl.Password)
	if err != nil {
		return err
	}

	return nil
}

func (i *Instance) GetEmployee(ctx context.Context, username string) (*models.Employee, error) {
	query := `SELECT id, username, password FROM employee WHERE username=$1`
	var user models.Employee
	err := i.Db.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
