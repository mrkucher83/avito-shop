// Package repository Работа с БД
package repository

type Authorization interface {
}

type Repository struct {
	Authorization
}

func NewService() *Repository {
	return &Repository{}
}
