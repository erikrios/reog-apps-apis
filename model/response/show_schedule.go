package response

import "time"

type ShowSchedule struct {
	ID       string `json:"id" extensions:"x-order=0"`
	GroupID  string
	Place    string
	StartOn  time.Time
	FinishOn time.Time
}
