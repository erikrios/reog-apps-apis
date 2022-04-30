package response

import "time"

type ShowSchedule struct {
	ID       string    `json:"id" extensions:"x-order=0"`
	GroupID  string    `json:"groupID" extensions:"x-order=0"`
	Place    string    `json:"place" extensions:"x-order=0"`
	StartOn  time.Time `json:"startOn" extensions:"x-order=0"`
	FinishOn time.Time `json:"finishOn" extensions:"x-order=0"`
}

type ShowScheduleDetails struct {
	ID       string    `json:"id" extensions:"x-order=0"`
	Group    Group     `json:"group" extensions:"x-order=1"`
	Place    string    `json:"place" extensions:"x-order=2"`
	StartOn  time.Time `json:"startOn" extensions:"x-order=3"`
	FinishOn time.Time `json:"finishOn" extensions:"x-order=4"`
}
