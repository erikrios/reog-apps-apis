package payload

type CreateGroup struct {
	Name      string `json:"name"`
	Leader    string `json:"leader"`
	Address   string `json:"address"`
	VillageID string `json:"villageId"`
}

type UpdateGroup struct {
	Name   string `json:"name"`
	Leader string `json:"leader"`
}
