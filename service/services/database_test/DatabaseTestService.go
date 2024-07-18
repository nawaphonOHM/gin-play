package service_database_test

type DatabaseTestService interface {
	GetData() *Response
}

type Response struct {
	Headers []string   `json:"headers"`
	Data    [][]string `json:"data"`
}
