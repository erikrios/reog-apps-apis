package address

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
)

type AddressRepository interface {
	Update(ctx context.Context, address entity.Address) (err error)
}
