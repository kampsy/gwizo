package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/google/uuid"
)

func (w walletService) Signin(id, password string) (string, error) {
	db := w.db

	formID := "phone_number"
	sql := "phone_number = ?"
	// user is trying to log in using their email address .
	if strings.Contains(id, "@") {
		formID = "email"
		sql = "email = ?"
	}

	var user Users
	// check if the username, phonenumber or email is in the database
	err := db.Where(sql, id).First(&user).Error
	if err != nil {
		return fmt.Sprintf("Wrong %s or password. Try again.", formID), err
	}

	// replace uuid with JWT
	token := uuid.New()
	tokenStr := token.String()

	return tokenStr, nil
}

var ErrSignin = errors.New("Unable to signin")
