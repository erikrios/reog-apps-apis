package admin

import (
	"context"

	"github.com/erikrios/reog-apps-apis/model/payload"
)

type AdminService interface {
	Login(ctx context.Context, credential payload.Credential) (token string, err error)
}
