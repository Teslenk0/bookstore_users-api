package mysql_utils

import (
	"fmt"
	"github.com/Teslenk0/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
	)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError{
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows){
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
		case 1062:
			return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError(fmt.Sprintf("error processing request: %s", err.Error()))
}