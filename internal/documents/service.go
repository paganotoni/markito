package documents

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"go.leapkit.dev/core/db"
)

type service struct {
	dbFn db.ConnFn
}

func (s service) sqlx() *sqlx.DB {
	conn, err := s.dbFn()
	if err != nil {
		return nil
	}

	return sqlx.NewDb(conn, "sqlite3")
}

func (s service) Save(id string, content string) error {
	db := s.sqlx()
	if db == nil {
		return errors.New("could not connect to database")
	}

	sql := `
		INSERT INTO
			documents (id, content)
		VALUES
			($1, $2)
		ON CONFLICT (id) DO UPDATE SET content = $2`
	_, err := db.Exec(sql, id, content)
	return err
}

func (s service) Find(id string) (*Document, error) {
	db := s.sqlx()
	if db == nil {
		return nil, errors.New("could not connect to database")
	}

	var doc Document
	err := db.Get(&doc, "SELECT * FROM documents WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (s service) List() ([]Document, error) {
	db := s.sqlx()
	if db == nil {
		return nil, errors.New("could not connect to database")
	}

	var doc []Document
	err := db.Select(&doc, "SELECT * FROM documents")
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func NewService(fn db.ConnFn) *service {
	return &service{
		dbFn: fn,
	}
}
