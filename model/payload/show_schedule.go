package payload

type CreateShowSchedule struct {
	GroupID string `json:"groupID" validate:"nonzero,min=2,max=10" extensions:"x-order=0"`
	Place   string `json:"place" validate:"nonzero,min=2,max=1000" extensions:"x-order=1"`
	// StartOn layout format: time.RFC822 (02 Jan 06 15:04 MST)
	StartOn string `json:"startOn" validate:"nonzero,min=2,max=30" extensions:"x-order=2"`
	// FinishOn layout format: time.RFC822 (02 Jan 06 15:04 MST)
	FinishOn string `json:"finishOn" validate:"nonzero,min=2,max=30" extensions:"x-order=3"`
}

type UpdateShowSchedule struct {
	Place string `json:"place" validate:"nonzero,min=2,max=1000" extensions:"x-order=0"`
	// StartOn layout format: time.RFC822 (02 Jan 06 15:04 MST)
	StartOn string `json:"startOn" validate:"nonzero,min=2,max=30" extensions:"x-order=1"`
	// FinishOn layout format: time.RFC822 (02 Jan 06 15:04 MST)
	FinishOn string `json:"finishOn" validate:"nonzero,min=2,max=30" extensions:"x-order=2"`
}
