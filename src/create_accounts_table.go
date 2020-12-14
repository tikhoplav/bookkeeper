package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"time"
)

func __main() {
	// Establish database connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	drop := flag.Bool("drop", false, "drop table before")
	flag.Parse()

	start := time.Now()

	sql := `CREATE TABLE IF NOT EXISTS "accounts" (
		"id" bigserial PRIMARY KEY,
		"code" text UNIQUE NOT NULL,
		"name" text NOT NULL,
		"desc" text NOT NULL DEFAULT '',
		"parent_id" bigint REFERENCES "accounts" ("id") ON DELETE CASCADE
	);`

	if *drop {
		prepend := `DROP TABLE IF EXISTS "accounts";`
		sql = fmt.Sprintf("%s\n%s", prepend, sql)
	}

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed executing migration: %v\n", err)
		os.Exit(1)
	}

	message := "\u001b[32mMigrated:\u001b[0m create_accounts_table (%v)\n"
	dt := time.Now().Sub(start).Truncate(time.Microsecond).String()
	fmt.Fprintf(os.Stderr, message, dt)
}
