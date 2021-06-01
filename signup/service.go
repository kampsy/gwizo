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


type Servicer interface{
	Signup(string,string,string,string, string, string,string, string, string, int, int, int, int) (string, error )
}

type Service struct{
	DB *gorm.DB
}

func ( svc Service) Signup(firstName,lastName, phone_number, password,pin,username, user_id,  email, balance , user_type_name string, status_id, account_type_id, account_id, user_type_id int ) ( string,  error ){
	db := svc.DB
	//UserID generating 
	cpuID := uuid.New()
	userID := fmt.Sprintf("us_%s", cpuID.String())
	//hashing the password 
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// hashing the pin
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.MinCost)
	if err != nil {
		return "" ,err
	}
	//generation of default colour 
	colour := database.RandomColour()
	//Generated value for the virtual account number 
	van := uuid.New()
	vanID := fmt.Sprintf("v.a.n_%s", van.String()) 


	user := database.Users{
		Model:                  gorm.Model{},
		UserID:                 userID,
		Firstname:              firstName,
		Lastname:               lastName,
		PhoneNumber:            phone_number,
		Password:               hashedPassword,
		Pin:                    hashedPin,
		Username:               username,
		AccountID:              account_id,
		Account:                database.Account{
			UserID: userID, 
			Balance: balance, 
			AccountTypeID: account_type_id, 
			StatusID: status_id,
		},
		UserTypeID:             user_type_id, // <<<<<======i thought of using the UserID instead of the  user_type_id but the data types in the structs are different
		UserType:               database.UserType{
			Name:	 user_type_name,
		},
		Colour:                 colour,
		LoginAttempts:          0,
		LoginTimeout:           time.Time{},
		Virtual_account_number: vanID, 
		Email:                  email,
	}
//Creating the User profile and Account 
 err = db.Create(&user).Error
	if err != nil{
		return "failed to create a user!!!!!!!!" ,err
	}

   var msg = "user successfully created!!!!!!"
	return  msg, nil 
}