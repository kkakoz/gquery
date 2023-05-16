package gquery

import "database/sql"

type query[T any] struct {
	db       *sql.DB
	sql      string
	vals     []any
	table    *Expr
	joins    []*Expr
	Selector []*Expr
}

func (s *query[T]) List() ([]*T, error) {
	return nil, nil
}

func (s *query[T]) One() (*T, error) {
	return nil, nil
}

func (s *query[T]) OneOrFail() (*T, error) {
	return nil, nil
}

func (s *query[T]) Delete() error {
	return nil
}

func (s *query[T]) Update(v any) *execResult {
	return newExecResult(nil, 0)
}

func (s *query[T]) UpdateColumn(c Column, v any) error {
	return nil
}

func (s *query[T]) ListRes() *result[[]*T] {
	return newResult([]*T{}, nil, 0)
}

func (s *query[T]) Where(exs ...*Expr) *query[T] {
	return s
}

func (s *query[T]) WhereIf(b bool, exs ...*Expr) *query[T] {
	if b {
		return s
	}
	return s
}

func (s *query[T]) WhereE(query string, args ...any) *query[T] {
	return s
}

func (s *query[T]) Table(table string) *query[T] {
	return s
}

func (s *query[T]) From(expr *Expr) *query[T] {
	return s
}

func (s *query[T]) WhereEx(query string, args ...any) *query[T] {
	return s
}

func (s *query[T]) Join(table string, exs ...*Expr) *query[T] {
	return s
}

func (s *query[T]) Select(fields ...any) *query[T] {
	return s
}

func Query[T any]() *query[T] {
	return &query[T]{}
}
