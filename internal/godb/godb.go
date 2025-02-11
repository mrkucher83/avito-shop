package godb

import "github.com/jackc/pgx/v4/pgxpool"

type Instance struct {
	Db *pgxpool.Pool
}

func (i *Instance) Close() {
	i.Db.Close()
}
