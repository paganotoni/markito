package documents

import (
	"github.com/jmoiron/sqlx"
)

type service struct {
	db *sqlx.DB
}

func (s service) Save(id string, content string) error {
	sql := `
		INSERT INTO 
			documents (id, content) 
		VALUES 
			($1, $2) 
		ON CONFLICT (id) DO UPDATE SET content = $2`
	_, err := s.db.Exec(sql, id, content)
	return err
}

func (s service) Find(id string) (*Document, error) {
	var doc Document
	err := s.db.Get(&doc, "SELECT * FROM documents WHERE id = $1", id)
	if err != nil {

		return nil, err
	}

	return &doc, nil
}

func NewService(db *sqlx.DB) *service {
	return &service{
		db: db,
	}
}
