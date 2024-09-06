package property

import "context"

type Storage interface {
	FindMany(ctx context.Context) ([]Property, error)
	FindOne(ctx context.Context, uuid string) (Property, error)
	Create(ctx context.Context, property Property) (string, error)
}
