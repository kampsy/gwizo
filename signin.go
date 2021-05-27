package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

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
	// check if the phonenumber or email is in the database
	err := db.Where(sql, id).First(&user).Error
	if err != nil {
		log.Println("Unable find user: ", err)
		return "", errors.New(fmt.Sprintf("Wrong %s or password. Try again.", formID))
	}

	// check the number of attempts
	if user.LoginAttempts == 3 {
		const longForm = "Jan 2, 2006 at 3:04pm (MST)"
		consTime, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (CAT)")

		currTime := time.Now()
		if user.LoginTimeout.Equal(consTime) {
			// New timeout
			err := db.Model(&Users{}).Where("user_id = ?", user.UserID).Update("login_timeout", currTime).Error
			if err != nil {
				return fmt.Sprintf("Wrong %s or password. Try again.", id), err
			}
		}

		// check the elapsed duration
		duration := currTime.Sub(user.LoginTimeout)
		fmt.Println(duration.Seconds())
		// If less than 30 seconds
		if duration.Seconds() < 30.0 {
			return "", errors.New(fmt.Sprintf("You attempted to login 3 times, please wait for %.0f seconds to elapse.", (30.0 - duration.Seconds())))
		}
		// Reset login attempt and timeout
		user.LoginAttempts = 0
		user.LoginTimeout = consTime
		err := db.Save(user).Error
		if err != nil {
			log.Println("Unable to update loginAttemps and LoginTimeout", err)
			return "", err
		}
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		log.Println("Password is not the same: ", err)

		// Update login attempts counter
		err := db.Model(&Users{}).Where("user_id = ?", user.UserID).Update("login_attempts", user.LoginAttempts+1).Error
		if err != nil {
			log.Print("Unable to update user login attemps: ", err)
			return "", errors.New(fmt.Sprintf("1-Wrong %s or password. Try again.", id))
		}

		return "", errors.New(fmt.Sprintf("2-Wrong %s or password. Try again.", id))
	}

	// Generate users unique token.
	tokenID := uuid.New()
	token := fmt.Sprintf("tok_%s", tokenID.String())

	tx := db.Begin()

	err = tx.Model(&Users{}).Where("user_id = ?", user.UserID).Update("token", token).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return fmt.Sprintf("Wrong %s or password. Try again.", id), err
	}

	tx.Commit()

	return token, nil
}

var ErrSignin = errors.New("Unable to signin")
var ErrThreeLoginAttempts = errors.New("3 login attemps, please try again after 30 seconds.")
