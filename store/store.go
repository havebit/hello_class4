package store

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type store struct {
	conn *pgx.Conn
}

func NewStore(conn *pgx.Conn) *store {
	return &store{conn: conn}
}

func (s *store) NewTask(title string) error {
	if _, err := s.conn.Exec(context.Background(), "INSERT INTO TODOS(title,done) VALUES($1, $2)", title, false); err != nil {
		return err
	}
	return nil
}
