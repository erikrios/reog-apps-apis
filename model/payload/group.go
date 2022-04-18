package payload

type CreateGroup struct {
	Name      string `json:"name" validate:"nonzero,min=2,max=80"`
	Leader    string `json:"leader" validate:"nonzero,min=2,max=80"`
	Address   string `json:"address" validate:"nonzero,min=2,max=1000"`
	VillageID string `json:"villageId" validate:"nonzero,min=2,max=20"`
}

type UpdateGroup struct {
	Name   string `json:"name" validate:"nonzero,min=2,max=80"`
	Leader string `json:"leader" validate:"nonzero,min=2,max=80"`
}
