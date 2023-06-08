package repository

import "context"

type IRepository[T any] interface {
	List(ctx context.Context) ([]T, error)
}
