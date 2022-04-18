package property

import (
	"context"

	"github.com/erikrios/reog-apps-apis/model/payload"
)

type PropertyService interface {
	Create(ctx context.Context, groupID string, p payload.CreateProperty) (id string, err error)
	Update(ctx context.Context, id string, p payload.UpdateProperty) (err error)
	Delete(ctx context.Context, id string) (err error)
	GeterateQRCode(ctx context.Context, id string) (file []byte, err error)
}
