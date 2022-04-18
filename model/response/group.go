package response

type Group struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Leader     string     `json:"leader"`
	Address    Address    `json:"address"`
	Properties []Property `json:"properties"`
}

type Address struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	VillageID    string `json:"villageId"`
	VillageName  string `json:"villageName"`
	DistrictID   string `json:"districtId"`
	DistrictName string `json:"districtName"`
	RegencyID    string `json:"regencyId"`
	RegencyName  string `json:"regencyName"`
	ProvinceID   string `json:"provinceId"`
	ProvinceName string `json:"provinceName"`
}

type Property struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      uint16 `json:"amount"`
}
