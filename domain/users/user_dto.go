package users

//This is the data transfer object, it moves between persistence and application
import (
	"github.com/Teslenk0/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	StatusActive = "active"
	StatusUsed= "used"
)

// User - This is the model
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
};

//Alias to a slice of users
type Users []User

//Validate - Function to validates user data
func (user *User) Validate() *rest_errors.RestError {

	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}

	if user.Password == "" {
		return rest_errors.NewBadRequestError("invalid password")
	}

	return nil
}
