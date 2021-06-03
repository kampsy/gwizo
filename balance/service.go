package balance

import (
	"dazwallet/database"

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

func (svc service) Balance(userid string) (string, error) {
	db := svc.db

	var account database.Account
	db.Where("user_id = ?", userid).Find(&account)
	if account.UserID == "" {
		return "", errTokenExpired
	}

	return account.Balance, nil
}
