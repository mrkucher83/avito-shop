// Package service бизнес логика
package service

import "github.com/mrkucher83/avito-shop/internal/repository"

type Authorization interface {
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
