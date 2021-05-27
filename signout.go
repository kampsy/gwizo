package main

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/google/uuid"
)

func (w walletService) Signout(token string) (string, error) {
	db := w.db

	if token == "" {
		return "", errors.New("token is empty mate")
	}

	// check if token exists
	var user Users
	db.Where("token = ?", token).Find(&user)
	if user.UserID == "" {
		return "", ErrSignout
	}

	// Generate users unique token.
	tokenID := uuid.New()
	newToken := fmt.Sprintf("tok_%s", tokenID.String())

	err := db.Model(&Users{}).Where("user_id = ?", user.UserID).Update("token", newToken).Error
	if err != nil {
		return "", err
	}

	return "Signout was successful", nil
}

var ErrSignout = errors.New("Unable to sign out")
