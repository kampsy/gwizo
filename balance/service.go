package balance

import (
	"dazwallet/database"

	"gorm.io/gorm"
)

type Servicer interface {
	Balance(string) (string, error)
}

// service ..
type service struct {
	DB *gorm.DB
}

func (svc service) Balance(userid string) (string, error) {
	db := svc.DB

	var account database.Account
	db.Where("user_id = ?", userid).Find(&account)

	return account.Balance, nil
}
