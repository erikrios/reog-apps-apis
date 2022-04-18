package group

import (
	"context"

	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/model/response"
)

type GroupService interface {
	Create(ctx context.Context, p payload.CreateGroup) (id string, err error)
	GetAll(ctx context.Context) (responses []response.Group, err error)
	GetByID(ctx context.Context, id string) (response response.Group, err error)
	Update(ctx context.Context, id string, p payload.UpdateGroup) (err error)
	Delete(ctx context.Context, id string) (err error)
	GeterateQRCode(ctx context.Context, id string) (file []byte, err error)
}
