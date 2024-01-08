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

func (n *Note) GetNoteById(id string) (*Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT
		id,
		title,
		description,
		content,
		created_at,
		updated_at FROM note WHERE id = $1
	`
	var note Note

	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
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

	return &note, nil
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

func (n *Note) UpdateNote(id string, body Note) (*Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `UPDATE note SET
		title = $1,
		description = $2,
		content = $3,
		updated_at = $4
		WHERE id = $5
		RETURNING *
	`
	_, err := db.ExecContext(
		ctx,
		query,
		body.Title,
		body.Description,
		body.Content,
		time.Now(),
		id,
	)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

func (n *Note) DeleteNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM note WHERE id = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
