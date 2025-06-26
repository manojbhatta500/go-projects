package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn

var Ctx context.Context

func ConnectToPostgress() bool {
	Ctx = context.Background()
	var err error
	Db, err = pgx.Connect(Ctx, os.Getenv("POSTGRESCONNECTIONSTRING"))
	if err != nil {
		fmt.Println("failed to connect to database")
		return false
	}

	fmt.Println("connected to database")
	return true
}

func CloseDbConnection() {
	Db.Close(Ctx)
}
