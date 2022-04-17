package property

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
)

type PropertyRepository interface {
	Insert(ctx context.Context, property entity.Property) (err error)
	Update(ctx context.Context, property entity.Property) (err error)
	Delete(ctx context.Context, id string) (err error)
}
