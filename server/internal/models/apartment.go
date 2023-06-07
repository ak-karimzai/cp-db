package models

type Apartment struct {
	ID          string     `json:"id"`
	Size        float32    `json:"size"`
	RoomNumbers int        `json:"room_numbers"`
	User        User       `json:"user"`
	Services    []*Service `json:"services,omitempty"`
}
