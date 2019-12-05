package services

import (
	"github.com/Teslenk0/bookstore_users-api/domain/users"
	"github.com/Teslenk0/bookstore_users-api/utils/errors"
)

//CreateUser - this function interacts with DB and creates a user, the error must be at the final of the return statement
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}
