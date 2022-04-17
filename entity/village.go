package entity

type Village struct {
	ID       string
	Name     string
	District District
}

type District struct {
	ID      string
	Name    string
	Regency Regency
}

type Regency struct {
	ID       string
	Name     string
	Province Province
}

type Province struct {
	ID   string
	Name string
}
