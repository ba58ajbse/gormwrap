// user_repository.go
package repository

import (
	"gormwrap/domain"
	"gormwrap/infrastructure/database"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	// Update(user *domain.User) error
	// Delete(id uint) error
	// FindAll() ([]domain.User, error)
}

type UserRepositoryGorm struct {
	db database.SQLHandler
}

func NewUserRepository(db database.SQLHandler) UserRepository {
	return &UserRepositoryGorm{db: db}
}

func (r *UserRepositoryGorm) Create(user *domain.User) error {
	return r.db.Create(&user)
}

func (r *UserRepositoryGorm) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id); err != nil {
		return nil, err
	}
	return &user, nil
}

// func (r *UserRepositoryGorm) Update(user *domain.User) error {
// 	return r.db.Save(user).Error
// }

// func (r *UserRepositoryGorm) Delete(id uint) error {
// 	return r.db.Delete(&domain.User{}, id).Error
// }

// func (r *UserRepositoryGorm) FindAll() ([]domain.User, error) {
// 	var users []domain.User
// 	if err := r.db.Find(&users).Error; err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }
