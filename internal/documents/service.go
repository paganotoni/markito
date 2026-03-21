package documents

import (
	"database/sql"
	"errors"

	"go.leapkit.dev/core/db"
	"markito/internal/markdown"
)

type Document struct {
	ID      string         `db:"id"`
	Name    sql.NullString `db:"name"`
	Content string         `db:"content"`
}

type service struct {
	dbFn db.ConnFn
}

func (s service) Save(id string, name string, content string) error {
	db, err := s.dbFn()
	if err != nil {
		return err
	}

	if db == nil {
		return errors.New("could not connect to database")
	}

	sql := `
		INSERT INTO
			documents (id, name, content)
		VALUES
			($1, $2, $3)
		ON CONFLICT (id) DO UPDATE SET name = $2, content = $3
	`

	_, err = db.Exec(sql, id, name, content)
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
	row := db.QueryRow("SELECT id, name, content FROM documents WHERE id = $1", id)
	err = row.Scan(&doc.ID, &doc.Name, &doc.Content)
	if err != nil {
		return nil, err
	}

	if !doc.Name.Valid || doc.Name.String == "" {
		doc.Name = sql.NullString{
			String: markdown.ExtractName(doc.Content),
			Valid:  true,
		}
	}

	return &doc, nil
}

func (s service) List() ([]Document, error) {
	db, err := s.dbFn()
	if err != nil {
		return nil, err
	}

	var doc []Document
	rows, err := db.Query("SELECT id, name, content FROM documents")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var d Document

		err := rows.Scan(&d.ID, &d.Name, &d.Content)
		if err != nil {
			return nil, err
		}

		if !d.Name.Valid || d.Name.String == "" {
			d.Name = sql.NullString{
				String: markdown.ExtractName(d.Content),
				Valid:  true,
			}
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
