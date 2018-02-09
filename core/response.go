package core

// Response Data model
type Response struct {
	// StateCode info
	StateCode int `json:"statecode"`
	// Message message
	Message string `json:"message"`
	// Data data
	Data interface{} `json:"data"`
}
