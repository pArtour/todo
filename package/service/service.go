package service

import "github.com/pArtour/todo/package/repository"

type Authorization interface {

}

type TodoList interface {

}

type TodoItem interface {
	
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}