package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var connStr = "postgres://postgres:manoj12@localhost:5432/expenses"

var Db *pgx.Conn

var Ctx context.Context

func ConnectToPostgress() bool {
	Ctx = context.Background()
	var err error
	Db, err = pgx.Connect(Ctx, connStr)
	if err != nil {
		fmt.Println("failed to connect to database")
		return false
	}

	// defer Db.Close(Ctx)
	fmt.Println("connected to database")
	return true
}

func CloseDbConnection() {
	Db.Close(Ctx)
}
