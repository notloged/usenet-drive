package sqllitequeue

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type SqlQueue interface {
	Enqueue(ctx context.Context, data string) error
	Dequeue(ctx context.Context, limit int) ([]Job, error)
	Delete(ctx context.Context, id int64) error
}

type Job struct {
	ID        int64
	Data      string
	CreatedAt time.Time
}

type sQLiteQueue struct {
	db *sql.DB
}

func NewSQLiteQueue(db *sql.DB) (SqlQueue, error) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS queue (
			id INTEGER PRIMARY KEY,
			data TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return nil, err
	}

	return &sQLiteQueue{db: db}, nil
}

func (q *sQLiteQueue) Enqueue(ctx context.Context, data string) error {
	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO queue (data) VALUES (?)")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, data)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (q *sQLiteQueue) Dequeue(ctx context.Context, limit int) ([]Job, error) {
	if limit < 1 {
		return nil, errors.New("limit must be greater than 0")
	}

	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, fmt.Sprintf("SELECT id, data, created_at FROM queue ORDER BY created_at ASC LIMIT %v", limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []Job
	for rows.Next() {
		var id int
		var data string
		var createdAt time.Time
		err = rows.Scan(&id, &data, &createdAt)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, Job{
			ID:        int64(id),
			Data:      data,
			CreatedAt: createdAt,
		})
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func (q *sQLiteQueue) Delete(ctx context.Context, id int64) error {
	tx, err := q.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM queue WHERE id = ?", id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
