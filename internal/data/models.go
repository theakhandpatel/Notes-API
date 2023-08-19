package data

import "database/sql"

type Models struct {
	Notes NoteModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Notes: NoteModel{DB: db},
	}
}
