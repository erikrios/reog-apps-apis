package response

type Group struct {
	ID         string     `json:"id" extensions:"x-order=0"`
	Name       string     `json:"name" extensions:"x-order=1"`
	Leader     string     `json:"leader" extensions:"x-order=2"`
	Address    Address    `json:"address" extensions:"x-order=3"`
	Properties []Property `json:"properties" extensions:"x-order=4"`
}

type Address struct {
	ID           string `json:"id" extensions:"x-order=0"`
	Address      string `json:"address" extensions:"x-order=1"`
	VillageID    string `json:"villageID" extensions:"x-order=2"`
	VillageName  string `json:"villageName" extensions:"x-order=3"`
	DistrictID   string `json:"districtID" extensions:"x-order=4"`
	DistrictName string `json:"districtName" extensions:"x-order=5"`
	RegencyID    string `json:"regencyID" extensions:"x-order=5"`
	RegencyName  string `json:"regencyName" extensions:"x-order=6"`
	ProvinceID   string `json:"provinceID" extensions:"x-order=7"`
	ProvinceName string `json:"provinceName" extensions:"x-order=8"`
}

type Property struct {
	ID          string `json:"id" extensions:"x-order=0"`
	Name        string `json:"name" extensions:"x-order=1"`
	Description string `json:"description" extensions:"x-order=2"`
	Amount      uint16 `json:"amount" extensions:"x-order=3"`
}
