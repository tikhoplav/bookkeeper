package main

import (
	"context"
	// "flag"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

type Account struct {
	id, parent_id    uint64
	code, name, desc string
}

func (a *Account) String() string {
	var format = "{\n\t\"id\": %d," +
		"\n\t\"code\": \"%s\"," +
		"\n\t\"name\": \"%s\"," +
		"\n\t\"desc\": \"%s\"," +
		"\n\t\"parent_id\": %d\n}"

	return fmt.Sprintf(format, a.id, a.code, a.name, a.desc, a.parent_id)
}

func main() {
	// Establish database connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	sql := `SELECT
			"id",
			"code",
			"name",
			"desc",
			coalesce("parent_id", 0) as "parent_id"
		FROM "accounts"`

	rows, err := conn.Query(context.Background(), sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to execute a query: %v\n", err)
		os.Exit(1)
	}

	var res []*Account

	for rows.Next() {
		var id, parent_id uint64
		var code, name, desc string
		err = rows.Scan(&id, &code, &name, &desc, &parent_id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading line: %v\n", err)
			return
		}
		n := &Account{id, parent_id, code, name, desc}
		res = append(res, n)
	}

	if rows.Err() != nil {
		fmt.Fprintf(os.Stderr, "Error executing a query: %v\n", err)
		os.Exit(1)
	}

	// Output result
	for _, item := range res {
		fmt.Fprintf(os.Stderr, "%v\n", item.String())
	}
}
