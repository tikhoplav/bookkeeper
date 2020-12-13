package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"time"
)

func abort(message string) {
	f := "\033[1;31mError:\u001b[0m %s\n"
	fmt.Fprintf(os.Stderr, f, message)
	os.Exit(1)
}

func main() {
	// Establish database connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	flag.Parse()
	var args = flag.Args()

	if len(args) < 2 {
		abort("at least 2 arguments required (code, name)")
	}

	code := args[0]
	name := args[1]

	desc := ""
	if len(args) > 2 {
		desc = args[2]
	}

	start := time.Now()
	sql := "INSERT INTO accounts(code, name, \"desc\") VALUES ($1, $2, $3)"
	_, err = conn.Exec(context.Background(), sql, code, name, desc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to execute a query: %v\n", err)
		os.Exit(1)
	}

	message := "\u001b[32mSuccess:\u001b[0m created (%v)\n"
	dt := time.Now().Sub(start).Truncate(time.Microsecond).String()
	fmt.Fprintf(os.Stderr, message, dt)
}