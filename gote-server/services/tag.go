package services

import "context"

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (t *Tag) CreateTag(tag Tag) (*Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO tag (name) 
		VALUES($1)
		RETURNING *`

	_, err := db.ExecContext(
		ctx,
		query,
		tag.Name,
	)

	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (t *Tag) GetAllTags() ([]*Tag, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name FROM tag`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var tags []*Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	return tags, nil
}
