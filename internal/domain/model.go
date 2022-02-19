package domain

type Weather struct {
	Temperature int     `json:"temperature"`
	Description *string `json:"description,omitempty"`
}
