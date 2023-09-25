package store

import "github.com/SharpDevOps10/GoPractice/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindById(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
