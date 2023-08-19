package data

import (
	"database/sql"
	"time"
)

type Note struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteModel struct {
	DB *sql.DB
}

func (nm NoteModel) GetAll() ([]*Note, error) {
	query := `
	SELECT id,title,body,created_at,updated_at 
	FROM notes`

	rows, err := nm.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes = []*Note{}
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.Id, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (nm NoteModel) Get(noteId int64) (*Note, error) {
	query := `
		SELECT title,body,created_at,updated_at 
		FROM notes
		WHERE id = $1`

	var note Note
	note.Id = noteId
	err := nm.DB.QueryRow(query, noteId).Scan(&note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (nm NoteModel) Insert(note *Note) error {
	query := `
		INSERT INTO notes (title,body,created_at,updated_at)
		VALUES ($1,$2,$3,$4)
		RETURNING id`

	err := nm.DB.QueryRow(query, note.Title, note.Body, note.CreatedAt, note.UpdatedAt).Scan(&note.Id)
	return err
}

func (nm NoteModel) Update(note *Note) error {
	query := `
		UPDATE notes
		SET title =$1, body=$2, updated_at=$3
		WHERE id=$4`

	_, err := nm.DB.Exec(query, note.Title, note.Body, note.UpdatedAt, note.Id)
	return err
}

func (nm NoteModel) Delete(noteId int64) error {
	query := `
		DELETE FROM notes
		WHERE id = $1`

	_, err := nm.DB.Exec(query, noteId)

	return err
}
