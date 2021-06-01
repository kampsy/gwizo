package database

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*Database tables and migrations
 */

// Users ...
type Users struct {
	gorm.Model
	UserID                 string `gorm:"type:varchar(255);not null" json:"userID"`
	Firstname              string `gorm:"type:varchar(20);not null" json:"firstname"`
	Lastname               string `gorm:"type:varchar(20);not null" json:"lastname"`
	PhoneNumber            string `gorm:"type:varchar(20);not null unique" json:"phone_number"`
	Password               []byte `gorm:"type:varbinary(255);not null" json:"password"`
	Pin                    []byte `gorm:"type:varbinary(255);not null unique" json:"pin"`
	Username               string `gorm:"type:varchar(255);not null unique" json:"username"`
	AccountID              int    `gorm:"type:int;not null unique" json:"account_id"`
	Account                Account
	UserTypeID             int `gorm:"type:int;not null unique" json:"user_type_id"`
	UserType               UserType
	Colour                 string `gorm:"type:varchar(255);not null unique" json:"color"`
	LoginAttempts          int
	LoginTimeout           time.Time `gorm:"type:time;not null unique" json:"login_timeout"`
	Virtual_account_number string    `gorm:"type:varchar(255);not null unique" json:"virtual_account_number"`
	Email                  string    `gorm:"type:varchar(50) not null unique" json:"email"`
}

// The role the user is playing. is it private user, admin etc
type UserType struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null unique" json:"user_type_name"`
}

type Account struct {
	gorm.Model
	UserID        string `gorm:"type:varchar(255);not null" json:"userID"`
	Balance       string `gorm:"type:varchar(255);not null unique" json:"amount"`
	AccountTypeID int    `gorm:"type:int;not null unique" json:"account_type_id"`
	AccountType   AccountType
	StatusID      int `gorm:"type:int;not null unique" json:"status_id"`
}

type AccountType struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null unique" json:"acount_type_name"`
}
type Status struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null unique" json:"status_name"`
}
type Evoucher struct {
	gorm.Model
	EvoucherID     string `gorm:"type:varchar(20);not null unique" json:"evoucher_number"`
	SenderNumber   string `gorm:"type:varchar(20);not null unique" json:"sender_number"`
	ReceiverNumber string `gorm:"type:varchar(20);not null unique" json:"receiver_number"`
	Amount         int    `gorm:"type:int;not null unique" json:"amount"`
}

type Transaction struct {
	gorm.Model
	UserID            string `gorm:"type:varchar(255);not null" json:"userID"`
	BalanceBefore     int    `gorm:"type:int;not null unique" json:"balance_before"`
	Amount            int    `gorm:"type:int;not null unique" json:"amount"`
	BalanceAfter      int    `gorm:"type:int;not null unique" json:"balance_after"`
	Note              string `gorm:"type:varchar(255);not null" json:"note"`
	StatusID          int    `gorm:"type:int;not null unique" json:"status_id"`
	Status            Status
	TransactionTypeID int `gorm:"type:int;not null unique" json:"transaction_type_id"`
	TransactionType   TransactionType
	TransferTypeID    int `gorm:"type:int;not null unique" json:"transfer_type_id"`
	TransferType      TransferType
}

type TransactionType struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null unique" json:"transaction_type_name"`
}
type TransferType struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null unique" json:"transfer_type_name"`
	Description string `gorm:"type:varchar(255);not null unique" json:"transfer_type_description"`
}

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

func randomColour() string {
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
	colour := randomColour()

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
	accountTypeName := "a"

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
