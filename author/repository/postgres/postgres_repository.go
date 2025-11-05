package postgres

import (
	"context"
	"database/sql"

	"github.com/grimoh/test-server/domain"
)

type postgresAuthorRepo struct {
	DB *sql.DB
}

func NewPostgresAuthorRepository(db *sql.DB) domain.AuthorRepository {
	return &postgresAuthorRepo{
		DB: db,
	}
}

func (m *postgresAuthorRepo) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Author, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.Author{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.Author{}

	err = row.Scan(
		&res.Id,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return
}

func (m *postgresAuthorRepo) GetById(ctx context.Context, id int64) (domain.Author, error) {
	query := `SELECT id, name, created_at, updated_at FROM author WHERE id=$1`
	return m.getOne(ctx, query, id)
}
