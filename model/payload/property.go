package payload

type CreateProperty struct {
	Name        string `json:"name" validate:"nonzero,min=2,max=80"`
	Description string `json:"description" validate:"nonzero,min=2,max=1000"`
	Amount      uint16 `json:"amount" validate:"nonzero,min=2"`
}

type UpdateProperty struct {
	Name        string `json:"name" validate:"nonzero,min=2,max=80"`
	Description string `json:"description" validate:"nonzero,min=2,max=1000"`
	Amount      uint16 `json:"amount" validate:"nonzero,min=2"`
}
