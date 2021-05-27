package main

import (
	"errors"

	"gorm.io/gorm"
)

/*Business logic
 */

// WalletService ...
type WalletService interface {
	Signin(string, string) (string, error)
	Signout(string) (string, error)
}

type walletService struct {
	db *gorm.DB
}

func (walletService) Signout(token string) (string, error) {
	// generate a new JWT and update the session in the database
	if token == "" {
		return "", errors.New("token is empty mate")
	}

	return "Signout was successful", nil
}

var ErrSignout = errors.New("Unable to sign out")
