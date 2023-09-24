package service

import (
	"Advertising/pkg/model"
	"Advertising/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
