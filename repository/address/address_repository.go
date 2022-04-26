package address

import (
	"context"

	"github.com/erikrios/reog-apps-apis/entity"
)

type AddressRepository interface {
	Update(ctx context.Context, id string, address entity.Address) (err error)
}
