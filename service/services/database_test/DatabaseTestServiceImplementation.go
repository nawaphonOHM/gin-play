package service_database_test

import repository_database_test "gin-play/repositories/database_test"

type DatabaseTestServiceImplementation struct {
	repository repository_database_test.DatabaseTestRepository
}

func (s *DatabaseTestServiceImplementation) GetData() *Response {
	response := s.repository.GetData()

	return &Response{
		Headers: response.Headers,
		Data:    response.Data,
	}
}

func NewDatabaseTestServiceImplementation(repository repository_database_test.DatabaseTestRepository) DatabaseTestService {
	return &DatabaseTestServiceImplementation{repository: repository}
}
