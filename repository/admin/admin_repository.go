package admin

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
)

type AdminRepository interface {
	FindByUsername(ctx context.Context, username string) (admin entity.Admin, err error)
}
