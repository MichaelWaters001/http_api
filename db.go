package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// create a new DB, trashing any existing ones with the same name
func initDB(ctx context.Context) (err error) {
	user := "michaelwaters"
	password := "api_pass"
	tableName := "api_test"

	//setup DB connection
	connStr := fmt.Sprintf("user=%s password=%s host=localhost port=5432 dbname=postgres pool_max_conns=10", user, password)
	apiTestDB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return
	}

	//clear existing tables
	dropTable := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	_, err = apiTestDB.Exec(ctx, dropTable)
	if err != nil {
		return
	}

	//setup empty table
	createTable := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id    VARCHAR,
		name  VARCHAR,
		email VARCHAR
	)`, tableName)

	_, err = apiTestDB.Exec(ctx, createTable)
	if err != nil {
		return
	}

	err = insertTestUsers(ctx)
	if err != nil {
		return
	}
	return
}

func insertTestUsers(ctx context.Context) (err error) {
	for _, user := range users {
		err = insertUser(user)
		if err != nil {
			fmt.Sprintf("Error inserting user:", err)
			return
		}
	}
	return
}

func insertUser(u user) error {
	query := "INSERT INTO api_test (id, name, email) VALUES ($1, $2, $3)"
	_, err := apiTestDB.Exec(ctx, query, u.ID, u.Name, u.Email)
	return err
}
