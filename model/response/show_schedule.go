package response

type ShowSchedule struct {
	ID      string `json:"id" extensions:"x-order=0"`
	GroupID string `json:"groupID" extensions:"x-order=1"`
	Place   string `json:"place" extensions:"x-order=2"`
	// StartOn layout format: time.RFC822 (02 Jan 06 15:04 MST)
	StartOn string `json:"startOn" extensions:"x-order=3"`
	// FinishOn layout format: time.RFC822 (02 Jan 06 15:04 MST)
	FinishOn string `json:"finishOn" extensions:"x-order=4"`
}

type ShowScheduleDetails struct {
	ID        string `json:"id" extensions:"x-order=0"`
	GroupID   string `json:"groupID" extensions:"x-order=1"`
	GroupName string `json:"groupName" extensions:"x-order=2"`
	Place     string `json:"place" extensions:"x-order=3"`
	// StartOn layout format: time.RFC822 (02 Jan 06 15:04 MST)
	StartOn string `json:"startOn" extensions:"x-order=4"`
	// FinishOn layout format: time.RFC822 (02 Jan 06 15:04 MST)
	FinishOn string `json:"finishOn" extensions:"x-order=5"`
}
