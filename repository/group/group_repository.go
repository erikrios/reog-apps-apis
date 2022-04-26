package group

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
)

type GroupRepository interface {
	Insert(ctx context.Context, group entity.Group) (err error)
	InsertAll(ctx context.Context, groups []entity.Group) (err error)
	FindAll(ctx context.Context) (groups []entity.Group, err error)
	FindByID(ctx context.Context, id string) (group entity.Group, err error)
	Update(ctx context.Context, id string, group entity.Group) (err error)
	Delete(ctx context.Context, id string) (err error)
}
