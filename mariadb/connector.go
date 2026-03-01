package mariadb

import (
	"database/sql"
	"fmt"
)

// Connector is a function that returns a database connection.
type Connector func() *sql.DB

func connector(host string, user string, pwd string) Connector {
	return func() *sql.DB {
		db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/?parseTime=true&multiStatements=true", user, pwd, host))

		return db
	}
}
