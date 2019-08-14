package permission

import (
	"context"
)

type Service interface {
	Set(ctx context.Context, identifier string, permission Permission) error
	Get(ctx context.Context, identifier string) ([]Permission, error)
	Has(ctx context.Context, identifier string, permission Permission) (bool, error)
}
