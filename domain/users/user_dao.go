package users

import (
	"fmt"
	"github.com/Teslenk0/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/Teslenk0/bookstore_users-api/utils/mysql_utils"
	"github.com/Teslenk0/bookstore_utils-go/logger"
	"github.com/Teslenk0/bookstore_utils-go/rest_errors"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

//Data access Object
//Interacts with DB

const (
	//QUERIES
	queryInsertUser                 = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES (?,?,?,?,?,?);"
	queryGetUser                    = "SELECT ID, first_name, last_name, email, date_created, status FROM users WHERE ID=?;"
	queryUpdateUser                 = "UPDATE users SET first_name=?, last_name=?, email=? WHERE ID=?;"
	queryDeleteUser                 = "DELETE FROM users WHERE ID=?;"
	queryFindUserByStatus           = "SELECT ID, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindUserByEmailAndPassword = "SELECT ID, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

//Get - Gets user by ID from DB - act like method
func (user *User) Get() *rest_errors.RestError {

	//Prepares the query
	stmt, err := users_db.Client.Prepare(queryGetUser)

	if err != nil {
		logger.Error("error when trying to prepare the get user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}

	//Close the stametent when the function returns
	defer stmt.Close()

	//Make a select and looks for only one result
	result := stmt.QueryRow(user.ID)

	//Populates the user given with the data from DB
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), "no rows in result set") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
		}

		logger.Error("error when trying to get user by id", getErr)
		return rest_errors.NewInternalServerError("database error", getErr)

	}

	return nil
}

//Save - saves object to DB - act like method
func (user *User) Save() *rest_errors.RestError {

	//Prepares the statement
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	//Ask if there was errors when attempting for preparing the stmt
	if err != nil {
		logger.Error("error when trying to prepare the save user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}
	//Close the connection when the functions returns
	defer stmt.Close()

	//Exec the statement
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when saving the user", saveErr)
		return rest_errors.NewInternalServerError("database error", saveErr)
	}

	//We get the last inserted id
	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when getting last inserted id after creating new user", err)
		return rest_errors.NewInternalServerError("database error", err)
	}

	user.ID = userId
	return nil
}

//Update - updates data from the database with the given one
func (user *User) Update() *rest_errors.RestError {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare the update user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("error when trying to update the user", err)
		return rest_errors.NewInternalServerError("database error",err)
	}
	return nil

}

//Delete - deletes a given user
func (user *User) Delete() *rest_errors.RestError {

	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare the delete user statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}

	defer stmt.Close()

	_, delErr := stmt.Exec(user.ID)
	if delErr != nil {
		logger.Error("error when trying to delete the user", delErr)
		return rest_errors.NewInternalServerError("database error", delErr)
	}
	return nil
}

//FindByStatus - find all the users with a given status
func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestError) {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare the search user statement", err)
		return nil, rest_errors.NewInternalServerError("database error", err)
	}

	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to search the users", err)
		return nil, rest_errors.NewInternalServerError("database error", err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying fill the struct with database data in search user method", err)
			return nil, rest_errors.NewInternalServerError("database error", err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError("no users matching given status")
	}

	return results, nil

}

//Login - find a user
func (user *User) FindByEmailAndPassword() *rest_errors.RestError {

	//Prepares the query
	stmt, err := users_db.Client.Prepare(queryFindUserByEmailAndPassword)

	if err != nil {
		logger.Error("error when trying to prepare the get user by email and password statement", err)
		return rest_errors.NewInternalServerError("database error", err)
	}

	//Close the stametent when the function returns
	defer stmt.Close()

	//Make a select and looks for only one result
	result := stmt.QueryRow(user.Email, user.Password, StatusActive)

	//Populates the user given with the data from DB
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return rest_errors.NewNotFoundError(fmt.Sprintf("user corresponding to %s not found", user.Email))
		}

		logger.Error("error when trying to get user by email and password", getErr)
		return rest_errors.NewInternalServerError("database error", getErr)

	}

	return nil
}
