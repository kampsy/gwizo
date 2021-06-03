// package transactions checks the most recent transactions
package transactions

import (
	"dazwallet/database"
	"fmt"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/leekchan/accounting"
	"gorm.io/gorm"
)

type Servicer interface {
	Trans(string) ([]data, error)
}

type service struct {
	db *gorm.DB
}

func (svc service) Trans(userID string) ([]data, error) {
	db := svc.db

	var transList []data

	ac := accounting.Accounting{Symbol: "K", Precision: 2}

	var tr []database.Transaction
	db.Limit(5).Where("sender = ?", userID).Or("receiver = ?", userID).Order("created_at desc").Find(&tr)

	for _, trans := range tr {
		if trans.Sender == userID {
			// signin user sent the money
			var userAcc database.Users
			db.Where("user_id = ?", trans.Receiver).Find(&userAcc)
			amountF64, err := strconv.ParseFloat(trans.Amount, 2)
			if err != nil {
				return transList, err
			}
			transList = append(transList, data{
				Type: "Money sent", Name: fmt.Sprintf("to %s", userAcc.Firstname), Note: trans.Note,
				Amount: " - " + ac.FormatMoney(amountF64), Time: humanize.Time(trans.CreatedAt),
			})

		} else {
			var userAcc database.Users
			db.Where("user_id = ?", trans.Receiver).Find(&userAcc)
			amountF64, err := strconv.ParseFloat(trans.Amount, 2)
			if err != nil {
				return transList, err
			}
			transList = append(transList, data{
				Type: "Money received", Name: fmt.Sprintf("from %s", userAcc.Firstname), Note: trans.Note,
				Amount: " - " + ac.FormatMoney(amountF64), Time: humanize.Time(trans.CreatedAt),
			})
		}
	}

	fmt.Println(tr)

	return transList, nil
}
