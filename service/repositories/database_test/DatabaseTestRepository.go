package repository_database_test

type DatabaseTestRepository interface {
	GetData() *Response
}

type Response struct {
	Headers []string
	Data    [][]string
}
