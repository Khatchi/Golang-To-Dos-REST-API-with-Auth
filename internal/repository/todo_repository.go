package repository

import (
	"context"
	"time"
	"todo_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// this is where we talk to our database
/*
	Milestone:
	Instead of passing *pgxpool.Pool directly into every function,
	a more idiomatic Go pattern is to use a repository struct
	in /repository, add===>
	type Reporitory struct{
		pool *pgxpool.Pool
	}
	Then in here,
	func (r *Repository) CreateTodo(title string, completed string) (*models.Todo, error){
		The same code as before...
	}

*/
func CreateTodo(pool *pgxpool.Pool, title string, completed bool) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO todos (title, completed)
		VALUES ($1, $2)
		RETURNING id, title, completed, created_at, updated_at
	`
	var todo models.Todo

	err := pool.QueryRow(ctx, query, title, completed).Scan(
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
