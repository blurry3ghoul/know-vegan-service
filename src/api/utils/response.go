package utils

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsePagination struct {
	Success  bool        `json:"success"`
	Metadata Metadata    `json:"metadata"`
	Data     interface{} `json:"data"`
}

type Metadata struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}
