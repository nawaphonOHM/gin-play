package repository_database_test

type DatabaseTestRepository interface {
	GetData() (*Response, error)
}

type Response struct {
	Headers []string
	Data    [][]string
}
