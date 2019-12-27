package users

import (
	"fmt"
	"github.com/Teslenk0/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/Teslenk0/bookstore_users-api/utils/errors"
	"github.com/Teslenk0/bookstore_users-api/utils/mysql_utils"
	_ "github.com/go-sql-driver/mysql"
)

//Data access Object
//Interacts with DB

const (
	//QUERIES
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	queryGetUser          = "SELECT ID, first_name, last_name, email, date_created FROM users WHERE ID=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE ID=?;"
	queryDeleteUser       = "DELETE FROM users WHERE ID=?;"
	queryFindUserByStatus = "SELECT ID, first_name, last_name, email, date_created, status FROM users WHERE ID=?;"
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

		return mysql_utils.ParseError(getErr)

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


	//Exec the statement
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	//We get the last inserted id
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.ID = userId

	return nil
}

//Update - updates data from the database with the given one
func (user *User) Update() *errors.RestError {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("internal server error: %s", err.Error()))
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil

}

//Delete - deletes a given user
func (user *User) Delete() *errors.RestError {

	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("internal server error: %s", err.Error()))
	}

	defer stmt.Close()

	_, delErr := stmt.Exec(user.ID)
	if delErr != nil {
		return mysql_utils.ParseError(delErr)
	}
	return nil
}

//FindByStatus - find all the users with a given status
func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("internal server error: %s", err.Error()))
	}

	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("internal server error: %s", err.Error()))
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError("no users matching given status")
	}

	return results, nil

}
