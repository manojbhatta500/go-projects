package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/manojbhatta500/auth/internal/config"
)

var Ctx context.Context
var Db *pgx.Conn

func ConnectToPostgres() bool {
	Ctx = context.Background()
	var err error
	Db, err = pgx.Connect(Ctx, config.ConfigInstance.GetDbConnectionString())
	if err != nil {
		fmt.Println("failed to connect to database:", err)
		return false
	}
	fmt.Println("connected to database")
	return true
}
