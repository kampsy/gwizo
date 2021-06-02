package database

import (
	"time"

	"gorm.io/gorm"
)

/*Database tables and migrations
 */

// Users ...
type Users struct {
	gorm.Model
	UserID                 string   `gorm:"type:varchar(255);not null" json:"userID"`
	Firstname              string   `gorm:"type:varchar(20);not null" json:"firstname"`
	Lastname               string   `gorm:"type:varchar(20);not null" json:"lastname"`
	PhoneNumber            string   `gorm:"type:varchar(20);not null unique" json:"phone_number"`
	Password               []byte   `gorm:"type:varbinary(255);not null" json:"password"`
	Pin                    []byte   `gorm:"type:varbinary(255);not null unique" json:"pin"`
	Username               string   `gorm:"type:varchar(255);not null unique" json:"username"`
	AccountID              int      `gorm:"type:int;not null unique" json:"account_id"`
	Account                Account  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserTypeID             int      `gorm:"type:int;not null unique" json:"user_type_id"`
	UserType               UserType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Colour                 string   `gorm:"type:varchar(255);not null unique" json:"color"`
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

type Transaction struct {
	gorm.Model
	Sender                string          `gorm:"type:varchar(255);not null" json:"sender"`   // can send money or a payment request request
	Receiver              string          `gorm:"type:varchar(255);not null" json:"receiver"` // can receive money or a payment request request
	Amount                string          `gorm:"type:varchar(255);not null unique" json:"amount"`
	SenderCharge          string          `gorm:"type:varchar(50);not null" json:"sender_charge"`
	SenderBalanceBefore   string          `gorm:"type:varchar(255);not null unique" json:"sender_balance_before"`
	SenderBalanceAfter    string          `gorm:"type:varchar(255);not null unique" json:"sender_balance_after"`
	ReceiverCharge        string          `gorm:"type:varchar(50);not null" json:"receiver_charge"` // just incase some countries charge users for receiving e money
	ReceiverBalanceBefore string          `gorm:"type:varchar(255);not null unique" json:"receiver_balance_before"`
	ReceiverBalanceAfter  string          `gorm:"type:varchar(255);not null unique" json:"receiver_balance_after"`
	Note                  string          `gorm:"type:varchar(255);not null" json:"note"`
	StatusID              int             `gorm:"type:int;not null unique" json:"status_id"`
	Status                Status          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransactionTypeID     int             `gorm:"type:int;not null unique" json:"transaction_type_id"`
	TransactionType       TransactionType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EvoucherID            int             `gorm:"type:int;not null unique" json:"evoucher_id"`
	Evoucher              Evoucher        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TransferTypeID        int             `gorm:"type:int;not null unique" json:"transfer_type_id"`
	TransferType          TransferType    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type TransactionType struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null unique" json:"transaction_type_name"`
}

type Evoucher struct {
	gorm.Model
	SenderNumber   string `gorm:"type:varchar(20);not null unique" json:"sender_number"`
	ReceiverNumber string `gorm:"type:varchar(20);not null unique" json:"receiver_number"`
	Amount         string `gorm:"type:varchar(255);not null unique" json:"amount"`
}

type TransferType struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255);not null unique" json:"transfer_type_name"`
	Description string `gorm:"type:varchar(255);not null unique" json:"transfer_type_description"`
}
