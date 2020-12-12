package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
	"os"
)

func main() {
	// Establish database connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	start := time.Now()

	sql := `CREATE EXTENSION IF NOT EXISTS ltree;
	CREATE TABLE IF NOT EXISTS accounts (
		"code" ltree,
		"name" TEXT UNIQUE NOT NULL,
		"desc" TEXT DEFAULT ''
	);
	CREATE INDEX code_gist_idx ON accounts USING gist(code);
	CREATE INDEX code_idx ON accounts USING btree(code);`

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed executing migration: %v\n", err)
		os.Exit(1)
	}

	message := "\u001b[32mMigrated:\u001b[0m create_accounts_table (%v)\n"
	dt := time.Now().Sub(start).Truncate(time.Microsecond).String()
	fmt.Fprintf(os.Stderr, message, dt)
}