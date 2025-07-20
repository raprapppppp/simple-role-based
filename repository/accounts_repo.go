package repository

import (
	"role-based/models"

	"gorm.io/gorm"
)

// Interface - Contains all available method in repository of Account
type AccountRepository interface {
	CreateAccountRepo(user models.Account) error
	CheckUsernameAlreadyExist(username string) bool
	LoginAccount(username string) (models.Account, error)
}

// DB Injection
type AccountDbRepo struct {
	db *gorm.DB
}

// To Initialize this Repository
func AccountRepositoryInit(db *gorm.DB) AccountRepository {
	return &AccountDbRepo{db}
}

func (r *AccountDbRepo) CreateAccountRepo(user models.Account) error {

	err := r.db.Create(&user).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *AccountDbRepo) CheckUsernameAlreadyExist(username string) bool {
	var user models.Account
	var count int64

	r.db.Model(&user).Where("Username = ?", username).Count(&count)

	if count > 0 {
		return true
	} else {
		return false
	}
}

func (r *AccountDbRepo) LoginAccount(username string) (models.Account, error) {
	var user models.Account

	err := r.db.Where("Username = ?", username).Find(&user).Error

	if err != nil {
		return models.Account{}, err
	}
	return user, nil
}
