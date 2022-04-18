package payload

type UpdateAddress struct {
	Address   string `json:"address"`
	VillageID string `json:"villageID"`
}
