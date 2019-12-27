package users

//This is the data transfer object, it moves between persistence and application
import (
	"strings"

	"github.com/Teslenk0/bookstore_users-api/utils/errors"
)

// User - This is the model
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate - Function to validates user data
func (user *User) Validate() *errors.RestError {

	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
