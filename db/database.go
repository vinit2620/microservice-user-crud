package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func ConnectDB() error {
	var err error
	Conn, err = pgx.Connect(context.Background(), "postgres://postgres:abcde@localhost:5432/user_service")
	if err != nil {
		return err
	}
	log.Println("Database connected successfully!")
	return nil
}
