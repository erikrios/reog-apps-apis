package payload

type UpdateAddress struct {
	Address   string `json:"address" validate:"nonzero,min=2,max=1000"`
	VillageID string `json:"villageID" validate:"nonzero,min=2,max=20"`
}
