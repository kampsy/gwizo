package signin

import (
	"dazwallet/auth"
	"dazwallet/database"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Service provides signin operations
type Servicer interface {
	Signin(string, string) (string, error)
}

type service struct {
	db *gorm.DB
}

func (svc service) Signin(username, password string) (string, error) {
	db := svc.db

	formID := "phone_number"
	sql := "phone_number = ?"
	// user is trying to log in using their email address .
	if strings.Contains(username, "@") {
		formID = "email"
		sql = "email = ?"
	}

	var user database.Users
	// check if the phonenumber or email is in the database
	err := db.Where(sql, username).First(&user).Error
	if err != nil {
		log.Println("Unable find user: ", err)
		return "", errors.New(fmt.Sprintf("Wrong %s or password. Try again.", formID))
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		log.Println("Password is not the same: ", err)

		// Update login attempts counter
		err := db.Model(&database.Users{}).Where("user_id = ?", user.UserID).Update("login_attempts", user.LoginAttempts+1).Error
		if err != nil {
			log.Print("Unable to update user login attemps: ", err)
			return "", errors.New(fmt.Sprintf("Wrong %s or password. Try again.", username))
		}

		return "", errors.New(fmt.Sprintf("Wrong %s or password. Try again.", username))
	}

	// jwt
	accessToken, err := auth.GenerateToken(user.UserID)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("Wrong %s or password. Try again.", username), err
	}
	fmt.Println(user.UserID)

	return accessToken, nil
}

var ErrSignin = errors.New("Unable to signin")
var ErrThreeLoginAttempts = errors.New("3 login attemps, please try again after 30 seconds.")
