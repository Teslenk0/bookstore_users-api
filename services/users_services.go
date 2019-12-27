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

	var result = &users.User{ID: userID}
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

//UpdateUser - this function updates the user with the data given
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {

	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}


	if isPartial {

		//Partial update
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		//Complete Update
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

//DeleteUser - this functions looks for a user and deletes it
func DeleteUser(userID int64) *errors.RestError{

	if userID <= 0 {
		return errors.NewBadRequestError("user identifier must be greater than 0")
	}

	var user = &users.User{ID: userID}
	return user.Delete()
}
