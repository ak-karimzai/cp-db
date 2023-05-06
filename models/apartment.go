package models

type Apartment struct {
	ID          string     `json:"id"`
	Size        float32    `json:"size"`
	RoomNumbers int        `json:"room_numbers"`
	UserId      string     `json:"user_id"`
	Services    []*Service `json:"services,omitempty"`
}
