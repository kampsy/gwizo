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

var errWrongUsernameOrPassword = errors.New("Wrong username or password")
var errGeneratingJWT = errors.New("error generating jwt")

func (svc service) Signin(username, password string) (string, error) {
	db := svc.db

	sql := "username = ?"
	// user is trying to log in using their email address .
	if strings.Contains(username, "@") {
		sql = "email = ?"
	}

	var user database.Users
	// check if the phonenumber or email is in the database
	err := db.Where(sql, username).First(&user).Error
	if err != nil {
		return "", errWrongUsernameOrPassword
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		// Update login attempts counter
		err := db.Model(&database.Users{}).Where("user_id = ?", user.UserID).Update("login_attempts", user.LoginAttempts+1).Error
		if err != nil {
			return "", errWrongUsernameOrPassword
		}

		return "", errWrongUsernameOrPassword
	}

	// jwt
	accessToken, err := auth.GenerateToken(user.UserID)
	if err != nil {
		log.Println(err)
		return "", errGeneratingJWT
	}
	fmt.Println(user.UserID)

	return accessToken, nil
}
