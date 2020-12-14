package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"strconv"
	"time"
)

func abort(message string) {
	f := "\033[1;31mError:\u001b[0m %s\n"
	fmt.Fprintf(os.Stderr, f, message)
	os.Exit(1)
}

func _main() {
	// Establish database connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	flag.Parse()
	var args = flag.Args()
	if len(args) < 3 {
		abort("at least 3 arguments required (parent_id, code, name)")
	}

	var input = []interface{}{
		nil,
		args[1],
		args[2],
		"",
	}

	parent_id, err := strconv.ParseInt(args[0], 0, 64)
	if err != nil {
		abort("parent_id must be type of int")
	}
	if parent_id > 0 {
		input[0] = parent_id
	}

	if len(args) > 3 {
		input[3] = args[3]
	}

	start := time.Now()
	var id uint64
	sql := `INSERT INTO accounts(
		"parent_id",
		"code",
		"name",
		"desc"
	) VALUES ($1, $2, $3, $4) RETURNING "id"`

	err = conn.QueryRow(context.Background(), sql, input...).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to execute a query: %v\n", err)
		os.Exit(1)
	}

	message := "\u001b[32mSuccess:\u001b[0m id: %d created (%v)\n"
	dt := time.Now().Sub(start).Truncate(time.Microsecond).String()
	fmt.Fprintf(os.Stderr, message, id, dt)
}
