package dto

type ApiResponse[T any] struct {
	Status int `json:"status"`
	Data   *T  `json:"data"`
}

type ApiResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
