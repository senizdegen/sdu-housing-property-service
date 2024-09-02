package property

import "context"

type Storage interface {
	FindMany(ctx context.Context) ([]Property, error)
}
