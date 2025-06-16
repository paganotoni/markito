package documents

import (
	"errors"

	"go.leapkit.dev/core/db"
)

type Document struct {
	ID      string `db:"id"`
	Content string `db:"content"`
}

type service struct {
	dbFn db.ConnFn
}

func (s service) Save(id string, content string) error {
	db, err := s.dbFn()
	if err != nil {
		return err
	}

	if db == nil {
		return errors.New("could not connect to database")
	}

	sql := `
		INSERT INTO
			documents (id, content)
		VALUES
			($1, $2)
		ON CONFLICT (id) DO UPDATE SET content = $2
	`

	_, err = db.Exec(sql, id, content)
	return err
}

func (s service) Find(id string) (*Document, error) {
	db, err := s.dbFn()
	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.New("could not connect to database")
	}

	var doc Document
	row := db.QueryRow("SELECT * FROM documents WHERE id = $1", id)
	err = row.Scan(&doc.ID, &doc.Content)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func (s service) List() ([]Document, error) {
	db, err := s.dbFn()
	if err != nil {
		return nil, err
	}

	var doc []Document
	rows, err := db.Query("SELECT * FROM documents")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d Document

		err := rows.Scan(&d.ID, &d.Content)
		if err != nil {
			return nil, err
		}

		doc = append(doc, d)
	}

	return doc, nil
}

func NewService(fn db.ConnFn) *service {
	return &service{
		dbFn: fn,
	}
}
