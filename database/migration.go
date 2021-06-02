package database

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Migrate
func Migrate(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(Users{}, Account{}, UserType{}, AccountType{}, Status{}, Evoucher{}, Transaction{}, TransactionType{}, TransferType{})

	// Has the test user been create
	var user Users
	db.Where("username = ?", "daz").Find(&user)
	if user.UserID == "" {
		createPrivateUser(db)
	}
}

func RandomColour() string {
	colours := []string{
		"red", "green", "brown", "orange", "purple", "pink",
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(colours) - 1)
	return colours[num]
}

// Create a private user for testing purposes
func createPrivateUser(db *gorm.DB) error {
	// Users
	colour := RandomColour()

	// Generate userID
	cpuID := uuid.New()
	userID := fmt.Sprintf("us_%s", cpuID.String())

	firstName := "daz"
	lastName := "wallet"
	username := "daz"

	phoneNumber := "+260967000000"

	password := "daz12345"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	pin := "12345"
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.MinCost)
	if err != nil {
		return err
	}

	// parse the time
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	loginTimeout, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (CAT)")

	// Generate virtual account number
	vanID := uuid.New()
	van := fmt.Sprintf("van_%s", vanID.String())

	email := "example@gmail.com"

	// UserType
	utn := "private-user"

	// Accounts
	balance := "10000.00"

	// AccountType
	accountTypeName := "private user"

	tx := db.Begin()

	// save to db
	err = tx.Create(&Users{
		UserID: userID, Firstname: firstName, Lastname: lastName, PhoneNumber: phoneNumber, Password: hashedPassword, Pin: hashedPin,
		Username: username, Account: Account{UserID: userID, Balance: balance, AccountType: AccountType{Name: accountTypeName}}, UserType: UserType{Name: utn},
		Colour: colour, LoginTimeout: loginTimeout, Virtual_account_number: van, Email: email,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
