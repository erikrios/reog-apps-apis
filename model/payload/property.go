package payload

type CreateProperty struct {
	Name        string `json:"name" validate:"nonzero,min=2,max=80" extensions:"x-order=0"`
	Description string `json:"description" validate:"nonzero,min=2,max=1000" extensions:"x-order=1"`
	Amount      uint16 `json:"amount" validate:"nonzero,min=1" extensions:"x-order=2"`
}

type UpdateProperty struct {
	Name        string `json:"name" validate:"nonzero,min=2,max=80" extensions:"x-order=0"`
	Description string `json:"description" validate:"nonzero,min=2,max=1000" extensions:"x-order=1"`
	Amount      uint16 `json:"amount" validate:"nonzero,min=1" extensions:"x-order=2"`
}
