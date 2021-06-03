package balance

import (
	"dazwallet/database"
	"strconv"

	"github.com/leekchan/accounting"
	"github.com/pkg/errors"

	"gorm.io/gorm"
)

type Servicer interface {
	Balance(string) (string, error)
}

// service ..
type service struct {
	db *gorm.DB
}

var errTokenExpired = errors.New("Token has expired")
var errUnableToFormatMoney = errors.New("Uable to format the money")

func (svc service) Balance(userid string) (string, error) {
	db := svc.db

	var account database.Account
	db.Where("user_id = ?", userid).Find(&account)
	if account.UserID == "" {
		return "", errTokenExpired
	}

	// Lets format the money.
	balanceF64, err := strconv.ParseFloat(account.Balance, 2)
	if err != nil {
		return "", errUnableToFormatMoney
	}

	ac := accounting.Accounting{Symbol: "K", Precision: 2}

	return ac.FormatMoney(balanceF64), nil
}
