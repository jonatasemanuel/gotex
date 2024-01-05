package services

import (
	"context"
	"time"
)

type Note struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (n *Note) GetAllNotes() ([]*Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT
		id,
		title,
		description,
		content,
		created_at,
		updated_at FROM note
	`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var notes []*Note
	for rows.Next() {
		var note Note
		err := rows.Scan(
			&note.ID,
			&note.Title,
			&note.Description,
			&note.Content,
			&note.CreatedAt,
			&note.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		notes = append(notes, &note)
	}

	return notes, nil
}

func (n *Note) CreateNote(note Note) (*Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO note
		(title, description, content, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5) RETURNING *
	`
	_, err := db.ExecContext(
		ctx,
		query,
		note.Title,
		note.Description,
		note.Content,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return &note, nil
}
