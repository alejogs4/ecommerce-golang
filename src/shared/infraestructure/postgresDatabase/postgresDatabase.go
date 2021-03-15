package postgresdatabase

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"
)

var dbOnce sync.Once
var database *sql.DB

// LoadDatabase open connection to database based in the passed parameters
func LoadDatabase(host, dbname, user string, port string) (*sql.DB, error) {
	var err error

	portNumber, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	dbOnce.Do(func() {
		database, err = loadDatabase(host, dbname, user, portNumber)
	})

	return database, err
}

func loadDatabase(host, dbname, user string, port int) (*sql.DB, error) {
	sqlStringConnection := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)

	database, err := sql.Open("postgres", sqlStringConnection)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return database, nil
}
