package usersdb

import "database/sql"

//This function is called every time the package is imported
func init() {
	usersDB, err := sql.Open(driverName, dataSourceName)
}
