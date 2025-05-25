package orm

import (
	"context"
	"database/sql"
)

// Querier Select
type Querier[T any] interface {
	Get(ctx context.Context) (*T, error)
	GetMulti(ctx context.Context) ([]*T, error)
}

// Executor Insert Delete Update
type Executor interface {
	Exec(ctx context.Context) (sql.Result, error)
}

// QueryBuilder builder
type QueryBuilder interface {
	Build() (*Query, error)
}

type Query struct {
	SQL  string
	Args []any
}
