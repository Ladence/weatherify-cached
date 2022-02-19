package weatherstack

type GetCurrentResponse struct {
	Request  interface{} `json:"request"`
	Location interface{} `json:"location"`
	Current  Current     `json:"current"`
}
