package village

import "github.com/erikrios/reog-apps-apis/entity"

type VillageRepository interface {
	FindByID(id string) (village entity.Village, err error)
}
