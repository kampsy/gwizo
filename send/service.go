// package send transfers money
package send

import (
	"dazwallet/database"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Servicer interface {
	Send(req sendData) (string, error)
}

type service struct {
	db *gorm.DB
}

var errNotNumber = errors.New("Amount not a number")
var errSenderNotFound = errors.New("sender not found")
var errSenderAccountNotFound = errors.New("sender account not found")
var errSenderBalanceDecimal = errors.New("Unable to parse sender balance")
var errSenderBalanceLessThanAmount = errors.New("Sender Balance is less than amount")
var errReceiverNotFound = errors.New("receiver not found")
var errReceiverAccountNotFound = errors.New("receiver account not found")
var errReceiverBalanceDecimal = errors.New("Unable to parse receiver balance")

// Send money to user
func (svc service) Send(req sendData) (string, error) {
	db := svc.db

	// is amount a number?
	amountF64, err := strconv.ParseFloat(req.amount, 2)
	if err != nil {
		return "", errNotNumber
	}

	// use big int
	amountDecimal := decimal.NewFromFloat(amountF64)

	// Get sender user account
	var sender database.Users
	db.Where("user_id = ?", req.sender).Find(&sender)
	if sender.UserID == "" {
		return "", errSenderNotFound
	}
	// Get sender account
	var senderAccount database.Account
	db.Where("user_id = ?", req.sender).Find(&senderAccount)
	if senderAccount.UserID == "" {
		return "", errSenderAccountNotFound
	}

	senderBalanceDecimal, err := decimal.NewFromString(senderAccount.Balance)
	if err != nil {
		return "", errSenderBalanceDecimal
	}
	senderBalanceBeforeF64, _ := senderBalanceDecimal.Float64()

	// Sender balance  <  the amount
	isSenderBalanceLess := senderBalanceDecimal.LessThan(amountDecimal)
	if isSenderBalanceLess {
		return "", errSenderBalanceLessThanAmount
	}

	// Deduct from sender balance
	senderBalanceAfterDecimal := senderBalanceDecimal.Sub(amountDecimal)
	senderBalanceAfterF64, _ := senderBalanceAfterDecimal.Float64()

	// Get receiver user account.
	var receiver database.Users
	db.Where("user_id = ?", req.receiver).Find(&receiver)
	if receiver.UserID == "" {
		return "", errReceiverNotFound
	}
	// Get receiver account
	var receiverAccount database.Account
	db.Where("user_id = ?", req.receiver).Find(&receiverAccount)
	if receiverAccount.UserID == "" {
		return "", errReceiverAccountNotFound
	}

	receiverBalanceDecimal, err := decimal.NewFromString(receiverAccount.Balance)
	if err != nil {
		return "", errReceiverBalanceDecimal
	}
	receiverBalanceBeforeF64, _ := receiverBalanceDecimal.Float64()

	receiverBalanceAfterDecimal := receiverBalanceDecimal.Add(amountDecimal)
	receiverBalanceAfterF64, _ := receiverBalanceAfterDecimal.Float64()

	// we are using float for decimal point formating.
	var transaction = database.Transaction{
		Sender:                req.sender,
		Receiver:              req.receiver,
		Amount:                fmt.Sprintf("%.2f", amountF64),
		SenderCharge:          "0.00",
		SenderBalanceBefore:   fmt.Sprintf("%.2f", senderBalanceBeforeF64),
		SenderBalanceAfter:    fmt.Sprintf("%.2f", senderBalanceAfterF64),
		ReceiverCharge:        "0.00",
		ReceiverBalanceBefore: fmt.Sprintf("%.2f", receiverBalanceBeforeF64),
		ReceiverBalanceAfter:  fmt.Sprintf("%.2f", receiverBalanceAfterF64),
		Note:                  req.note,
		TransactionType:       database.TransactionType{Name: "send"},
		Status:                database.Status{Name: "completed"},
		Evoucher:              database.Evoucher{Amount: "0.00"},
		TransferType: database.TransferType{
			Name:        "transfer",
			Description: "moved money from one account to another",
		},
	}

	// Start transaction
	tx := db.Begin()

	// Create a transaction
	err = tx.Save(&transaction).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	// update sender balance
	senderAccount.Balance = fmt.Sprintf("%.2f", senderBalanceAfterF64)
	sender.Account = senderAccount
	err = tx.Save(&senderAccount).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	// update receiver balance
	receiverAccount.Balance = fmt.Sprintf("%.2f", receiverBalanceAfterF64)
	receiver.Account = receiverAccount
	err = tx.Save(&receiverAccount).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()

	return "money sent!!", nil
}
