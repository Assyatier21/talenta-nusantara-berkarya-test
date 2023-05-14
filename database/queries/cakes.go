package database

const (
	GetCakes = `SELECT id, title, description, rating, image, created_at, updated_at, deleted_at 
					FROM cakes WHERE deleted_at IS NULL ORDER BY rating DESC, title ASC 
					LIMIT ? OFFSET ?`

	GetCakeDetails = `SELECT id, title, description, rating, image, created_at, updated_at, deleted_at 
						FROM cakes WHERE id = ? AND deleted_at IS NULL`

	InsertCake = `INSERT INTO cakes (title, description, rating, image, created_at, updated_at) 
                    VALUES (?, ?, ?, ?, ?, ?)`

	UpdateCake = `UPDATE cakes SET title = ?, description = ?, rating = ?, image = ?, updated_at = ?
                    WHERE id = ? AND deleted_at IS NULL`

	DeleteCake = `UPDATE cakes SET deleted_at = ?
					WHERE id = ?`
)
