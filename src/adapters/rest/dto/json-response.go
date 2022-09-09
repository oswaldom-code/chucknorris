package dto

type MessageResponse struct {
	Status  string `json:"status"`
	Message string `json:"mensaje"`
}

type ResponseWithData struct {
	Status  string      `json:"status"`
	Message string      `json:"mensaje"`
	Data    interface{} `json:"data"`
}
