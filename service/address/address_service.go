package address

import (
	"context"

	"github.com/erikrios/reog-apps-apis/model/payload"
)

type AddressService interface {
	Update(ctx context.Context, id string, p payload.UpdateAddress) (err error)
}
