package main

import (
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
