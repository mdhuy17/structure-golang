package repository

import (
	"PresentationProject/internal/model"
	"gorm.io/gorm"
)

type InterfaceUserRepository interface {
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(email string) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user *model.User

	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Update(user *model.User) (*model.User, error) {
	if err := u.db.Where("email = ?", user.Email).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Delete(email string) error {
	if err := u.db.Where("email = ?", email).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
