package main

import (
	"time"
	"github.com/jinzhu/gorm"
)

/*Database tables and migrations
 */

// Users ...
type Users struct {
	gorm.Model
	UserID string 		`gorm:"type:varchar(20);not null" json:"userID"`
	Firstname string 	`gorm:"type:varchar(20);not null" json:"firstname"`
	Lastname string 	`gorm:"type:varchar(20);not null" json:"lastname"`
	PhoneNumber string  `gorm:"type:varchar(20);not null unique" json:"phone_number"`
	Password []byte		`gorm:"type:varbinary(50);not null" json:"password"`
	Pin 	 []byte 	`gorm:"type:varbinary(20);not null unique" json:"pin"`
	Username string 	`gorm:"type:varchar(255);not null unique" json:"username"`
	UserTypeID int 		`gorm:"type:int;not null unique" json:"user_type_id"`
	StatusID int 		`gorm:"type:int;not null unique" json:"status_id"`
	Colour string 		`gorm:"type:varchar(255);not null unique" json:"color"`
	Login_timeout time.Time `gorm:"type:time;not null unique" json:"login_timeout"`
	Virtual_account_number string `gorm:"type:varchar(255);not null unique" json:"virtual_account_number"`
	Avatar string 				`gorm:"type:varchar(255) not null" json:"avatar"`
	Notification_token string 	`gorm:"type:varchar(50) not null unique" json:"notification"`
	Email string 				`gorm:"type:varchar(50) not null unique" json:"email"`
}

type Account struct {
	gorm.Model
	UserID string `gorm:"type:varchar(20);not null" json:"userID"`
	Balance int 	`gorm:"type:int;not null unique" json:"amount"`
	AccountTypeID int `gorm:"type:int;not null unique" json:"account_type_id"`
	StatusID int `gorm:"type:int;not null unique" json:"status_id"`
}

type UserType struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null unique" json:"user_type_name"`
}
type AccountType struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null unique" json:"acount_type_name"`
}
type Status struct{
	gorm.Model
	Name string	`gorm:"type:varchar(255);not null unique" json:"status_name"`
}
type Evoucher struct{
	gorm.Model
	EvoucherID string	`gorm:"type:varchar(20);not null unique" json:"evoucher_number"`
	SenderNumber string	 `gorm:"type:varchar(20);not null unique" json:"sender's_number"`
	ReceiverNumber string `gorm:"type:varchar(20);not null unique" json:"receiver's_number"`
	Amount int 			`gorm:"type:int;not null unique" json:"amount"`
}

type Transaction struct{
	gorm.Model
	UserID string `gorm:"type:varchar(20);not null" json:"userID"`
	BalanceBefore  int `gorm:"type:int;not null unique" json:"balance_before"`
	Amount int `gorm:"type:int;not null unique" json:"amount"`
	BalanceAfter int `gorm:"type:int;not null unique" json:"balance_after"`
	StatusID int `gorm:"type:int;not null unique" json:"status_id"`
	TransactionTypeID int `gorm:"type:int;not null unique" json:"transaction_type_id"`
	TransferTypeID int `gorm:"type:int;not null unique" json:"transfer_type_id"`
}

type transactionType struct{
	gorm.Model
	Name string `gorm:"type:varchar(255);not null unique" json:"transaction_type_name"`
}
type TransferType struct{
	gorm.Model
	Name string `gorm:"type:varchar(255);not null unique" json:"transfer_type_name"`
	Description string `gorm:"type:varchar(255);not null unique" json:"transfer_type_description"`
}


