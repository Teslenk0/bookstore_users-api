package services

//Bussiness logic of our users
import (
	"github.com/Teslenk0/bookstore_users-api/domain/users"
	"github.com/Teslenk0/bookstore_users-api/utils/errors"
)

//GetUser - this function interacts with DB and gets a user
func GetUser(userID int64) (*users.User, *errors.RestError) {

	if userID <= 0 {
		return nil, errors.NewBadRequestError("User Identifier Must be Greater than 0")
	}

	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

//CreateUser - this function interacts with DAO and creates a user, the error must be at the final of the return statement
func CreateUser(user users.User) (*users.User, *errors.RestError) {

	//User validates itself
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
