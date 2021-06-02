// package search adds search functionality to the wallet. search for people within dazwallet
package search

import (
	"dazwallet/database"

	"gorm.io/gorm"
)

// Servicer our search service interface
type Servicer interface {
	Search(string) ([]data, error)
}

type service struct {
	db *gorm.DB
}

// Search search for private users in the database
func (svc service) Search(query string) ([]data, error) {
	db := svc.db

	var userList []database.Users
	db.Where("firstname LIKE ?", "%"+query+"%").Find(&userList)

	var searchResult []data
	for _, user := range userList {
		searchResult = append(searchResult, data{
			FirstName: user.Firstname, LastName: user.Lastname, UserName: user.Username, PhoneNumber: user.PhoneNumber,
			Email: user.Email, UserID: user.UserID,
		})
	}

	return searchResult, nil
}
