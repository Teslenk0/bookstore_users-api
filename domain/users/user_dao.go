package users

import (
	"fmt"

	"github.com/Teslenk0/bookstore_users-api/utils/date"
	"github.com/Teslenk0/bookstore_users-api/utils/errors"
)

//Data access Object
//Interacts with DB

var (
	usersDB = make(map[int64]*User)
)

//Get - Gets user by ID from DB - act like method
func (user *User) Get() *errors.RestError {

	result := usersDB[user.ID]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	//There is no erros
	return nil
}

//Save - saves object to DB - act like method
func (user *User) Save() *errors.RestError {

	current := usersDB[user.ID]

	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError("User " + user.Email + " Already Registered")
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d Already Exists", user.ID))
	}

	user.DateCreated = date.GetNowString()
	usersDB[user.ID] = user
	//Retun nil because there is no error
	return nil
}
