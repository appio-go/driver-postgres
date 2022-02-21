package driver_postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func Connection() (*sql.DB, error) {
	postgresConn, _ := os.LookupEnv("POSTGRES_CONNECTION")
	db, err := sql.Open("postgres", postgresConn)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
