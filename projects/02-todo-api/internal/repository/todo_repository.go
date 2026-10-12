package repository

import (
	"context"
	"time"
	"todo-api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// function to create todos
func CreateTodo(pool *pgxpool.Pool, title string, completed bool) (*models.Todo, error) {
	// timeout context
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	// SQL query
	var query string = `
		INSERT INTO todos(title, copmleted)
		VALUES ($1, $2)
		RETURNING id, title, completed, created_at, updated_at
	`

	// query execution and scanning the result
	var todo models.Todo
	var err error = pool.QueryRow(ctx, query, title, completed).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
