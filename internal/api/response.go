package api

type APIResponse struct {
	Status bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
