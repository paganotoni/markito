package documents

type Document struct {
	ID      string `db:"id"`
	Content string `db:"content"`
}
