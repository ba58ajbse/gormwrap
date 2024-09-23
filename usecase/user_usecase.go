// user_usecase.go
package usecase

import (
	"gormwrap/domain"
	"gormwrap/infrastructure/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	return u.userRepo.Create(user)
}

// 他のメソッドも同様に実装
