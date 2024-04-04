package Model

import "encoding/json"

type Reservation struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Time  string `json:"time"`
}

func (r *Reservation) MarshalBinary() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Reservation) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, r)
}
