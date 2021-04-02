package models

type Event struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"owner"`
}

func (e Event) Create() bool {
	return true
}

func (e Event) Delete() bool {
	return true
}
