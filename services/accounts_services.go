package services

import (
	"fmt"
	"role-based/config/encrypt"
	"role-based/models"
	"role-based/repository"
)

// Interface - Contains all available method in Services of Account
type AccountServices interface {
	CreateAccountService(user models.Account) (string, error)
	AccountLoginService(user models.LoginCred) (models.Account, string)
	GetProfileService(id int) (models.Account, error)
}

// Repository Injection
type AccountRepoInjection struct {
	repo repository.AccountRepository
}

// To Initialize this Services
func AccountServicesInit(rep repository.AccountRepository) AccountServices {
	return &AccountRepoInjection{rep}
}

func (s *AccountRepoInjection) CreateAccountService(user models.Account) (string, error) {

	isExist := s.repo.CheckUsernameAlreadyExist(user.Username)
	if !isExist {
		user.Password = encrypt.HashPassword(user.Password)
		fmt.Print(user)

		err := s.repo.CreateAccountRepo(user)
		if err != nil {
			return "", err
		}
		return "Created", nil

	}
	return "Already Exist", nil
}

func (s *AccountRepoInjection) AccountLoginService(user models.LoginCred) (models.Account, string) {

	isAlreadyExist := s.repo.CheckUsernameAlreadyExist(user.Username)
	if !isAlreadyExist {
		return models.Account{}, "User Does not exist exist"
	}

	userAccount, err := s.repo.LoginAccount(user.Username)
	if err != nil {
		return models.Account{}, "Error in Database"
	}

	isMatch := encrypt.CompareHashAndPassword(userAccount.Password, user.Password)
	if !isMatch {
		return models.Account{}, "Password does not match"
	}
	return userAccount, "Account match"
}

func (s *AccountRepoInjection) GetProfileService(id int) (models.Account, error) {
	return s.repo.GetProfile(id)
}
