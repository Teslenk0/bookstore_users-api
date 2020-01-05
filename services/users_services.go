package services

//Bussiness logic of our users
import (
	"github.com/Teslenk0/bookstore_users-api/domain/users"
	"github.com/Teslenk0/bookstore_users-api/utils/crypto_utils"
	"github.com/Teslenk0/bookstore_users-api/utils/date"
	"github.com/Teslenk0/bookstore_users-api/utils/errors"
)

//Interface with methods
type usersServiceInterface interface {
	GetUser(userID int64) (*users.User, *errors.RestError)
	CreateUser(user users.User) (*users.User, *errors.RestError)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError)
	DeleteUser(userID int64) *errors.RestError
	SearchUser(status string) (users.Users, *errors.RestError)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestError)
}

//Struct
type usersService struct {
}

//Implementing the interface
var (
	UsersService usersServiceInterface = &usersService{}
)

//GetUser - this function interacts with DB and gets a user
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestError) {

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
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {

	//User validates itself
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date.GetNowDBString()
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

//UpdateUser - this function updates the user with the data given
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {

	current, err := UsersService.GetUser(user.ID)
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
func (s *usersService) DeleteUser(userID int64) *errors.RestError {

	if userID <= 0 {
		return errors.NewBadRequestError("user identifier must be greater than 0")
	}

	var user = &users.User{ID: userID}
	return user.Delete()
}

//FindByStatus - this functions asks the dao for users with the given status
func (s *usersService) SearchUser(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	encryptedPassword := crypto_utils.GetMd5(request.Password)
	dao := &users.User{
		Email:    request.Email,
		Password: encryptedPassword,
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
