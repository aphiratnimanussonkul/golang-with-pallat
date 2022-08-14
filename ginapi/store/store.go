package store

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:mysecretpassword@localhost:5432/myapp")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

type store struct {
	conn *pgx.Conn
}

func NewStore(conn *pgx.Conn) *store {
	return &store{conn: conn}
}

func (s *store) NewTask(title string) error {
	if _, err := s.conn.Exec(context.Background(), "INSERT INTO TODOS(TITLE) VALUES($1)", title); err != nil {
		return err
	}
	return nil
}
