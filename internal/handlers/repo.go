package handlers

import "github.com/mrkucher83/avito-shop/internal/godb"

type Repo struct {
	storage *godb.Instance
}

func NewRepo(r *godb.Instance) *Repo {
	return &Repo{storage: r}
}
