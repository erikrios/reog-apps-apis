package payload

type CreateGroup struct {
	Name      string `json:"name" validate:"nonzero,min=2,max=80" extensions:"x-order=0"`
	Leader    string `json:"leader" validate:"nonzero,min=2,max=80" extensions:"x-order=1"`
	Address   string `json:"address" validate:"nonzero,min=2,max=1000" extensions:"x-order=2"`
	VillageID string `json:"villageID" validate:"nonzero,min=2,max=20" extensions:"x-order=3"`
}

type UpdateGroup struct {
	Name   string `json:"name" validate:"nonzero,min=2,max=80" extensions:"x-order=0"`
	Leader string `json:"leader" validate:"nonzero,min=2,max=80" extensions:"x-order=1"`
}
