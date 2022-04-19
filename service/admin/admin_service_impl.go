package admin

import (
	"context"

	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository/admin"
	"github.com/erikrios/reog-apps-apis/service"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	"gopkg.in/validator.v2"
)

type adminServiceImpl struct {
	adminRepository   admin.AdminRepository
	passwordGenerator generator.PasswordGenerator
	tokenGenerator    generator.TokenGenerator
}

func NewAdminServiceImpl(
	adminRepository admin.AdminRepository,
	passwordGenerator generator.PasswordGenerator,
	tokenGenerator generator.TokenGenerator,
) *adminServiceImpl {
	return &adminServiceImpl{
		adminRepository:   adminRepository,
		passwordGenerator: passwordGenerator,
		tokenGenerator:    tokenGenerator,
	}
}

func (a *adminServiceImpl) Login(ctx context.Context, credential payload.Credential) (token string, err error) {
	if validateErr := validator.Validate(credential); validateErr != nil {
		err = service.ErrInvalidPayload
		return
	}

	admin, repoErr := a.adminRepository.FindByUsername(ctx, credential.Username)
	if repoErr != nil {
		err = service.MapError(repoErr)
		return
	}

	if compareErr := a.passwordGenerator.CompareHashAndPassword([]byte(admin.Password), []byte(credential.Password)); compareErr != nil {
		err = service.ErrCredentialNotMatch
		return
	}

	token, genErr := a.tokenGenerator.GenerateToken(admin.ID, admin.Username)
	if genErr != nil {
		err = service.MapError(genErr)
	}
	return
}
