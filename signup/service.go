package signup

import (
	// "dazwallet/auth"
	"dazwallet/database"
	"fmt"

	// "log"
	// "strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Servicer interface {
	Signup(req signupData) (string, error)
}

type service struct {
	DB *gorm.DB
}

func (svc service) Signup(data signupData) (string, error) {
	db := svc.DB

	//UserID generating
	cpuID := uuid.New()
	userID := fmt.Sprintf("us_%s", cpuID.String())

	//hashing the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// hashing the pin
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(data.Pin), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	loginTimeout, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (CAT)")

	//generation of default colour
	colour := database.RandomColour()

	//Generated value for the virtual account number
	van := uuid.New()
	vanID := fmt.Sprintf("v.a.n_%s", van.String())

	user := database.Users{
		UserID:      userID,
		Firstname:   data.FirstName,
		Lastname:    data.LastName,
		PhoneNumber: data.PhoneNumber,
		Password:    hashedPassword,
		Pin:         hashedPin,
		Username:    data.Username,
		Account: database.Account{
			UserID:  userID,
			Balance: data.Balance,
			AccountType: database.AccountType{
				Name: data.AccountTypeName,
			},
		},
		UserType: database.UserType{
			Name: data.UserTypeName,
		},
		Colour:                 colour,
		LoginAttempts:          0,
		LoginTimeout:           loginTimeout,
		Virtual_account_number: vanID,
		Email:                  data.Email,
	}

	//Creating the User profile and Account
	err = db.Create(&user).Error
	if err != nil {
		return "failed to create a user!!!!!!!!", err
	}

	var msg = "user successfully created!!!!!!"
	return msg, nil
}
