package dto

type WebStandardResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}
