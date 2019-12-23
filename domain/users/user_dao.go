package users

import (
	"fmt"
	"github.com/Teslenk0/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/Teslenk0/bookstore_users-api/utils/date"
	"github.com/Teslenk0/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

//Data access Object
//Interacts with DB

const (
	//ERROR HANDLING
	emptyResultSet   = "no rows in result set"
	indexUniqueEmail = "for key 'users_email_uindex'"

	//QUERIES
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?,?,?,?);"
	queryGetUser    = "SELECT * FROM users WHERE ID=?;"
)

var (
	usersDB = make(map[int64]*User)
)

//Get - Gets user by ID from DB - act like method
func (user *User) Get() *errors.RestError {

	//Prepares the query
	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error getting the user: %s", err.Error()))
	}

	//Close the stametent when the function returns
	defer stmt.Close()

	//Make a select and looks for only one result
	result := stmt.QueryRow(user.ID)

	//Populates the user given with the data from DB
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {

		//Asks if result is empty
		if strings.Contains(getErr.Error(), emptyResultSet) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d does not exists", user.ID))
		}

		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user: %s", getErr.Error()))
	}

	return nil
}

//Save - saves object to DB - act like method
func (user *User) Save() *errors.RestError {

	//Prepares the statement
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	//Ask if there was errors when attempting for preparing the stmt
	if err != nil {
		return errors.NewInternalServerError("error when tying to save user")
	}
	//Close the connection when the functions returns
	defer stmt.Close()

	//Set the date as now
	user.DateCreated = date.GetNowString()

	//Exec the statement
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {

		//Attempts to convert to mysql error
		sqlErr, ok := saveErr.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError("error when attempting to convert to mysql error")
		}

		//Check error by number
		switch sqlErr.Number {
		case 1062:
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))

		default:
			return errors.NewInternalServerError(fmt.Sprintf("error when tying to save user: %s", saveErr.Error()))
		}
	}

	//We get the last inserted id
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when tying to save user: %s", err.Error()))
	}

	user.ID = userId

	return nil
}
