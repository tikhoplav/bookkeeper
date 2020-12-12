package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4"
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

	// Parse input, define the procedure
	n := flag.Int("n", 10, "number of items in series")
	flag.Parse()

	// Start process, generate the result
	var res []int32
	fmt.Fprintf(os.Stderr, "Starting test with n: %v\n", *n)

	rows, err := conn.Query(context.Background(), "select generate_series(1,$1)", *n)
	if err != nil {
	    fmt.Fprintf(os.Stderr, "Unable to execute a query%v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
	    var n int32
	    err = rows.Scan(&n)
	    if err != nil {
	        return
	    }
	    res = append(res, n)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "Error executing a query%v\n", err)
		os.Exit(1)
	}

	// Output result
	for _, item := range res {
		fmt.Fprintf(os.Stderr, "%v\n", item)
	}
}