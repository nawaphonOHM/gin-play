package service_database_test

import repository_database_test "gin-play/repositories/database_test"

type DatabaseTestServiceImplementation struct {
	repository repository_database_test.DatabaseTestRepository
}

func (s *DatabaseTestServiceImplementation) GetData() *Response {
	//TODO implement me
	panic("implement me")
}

func NewDatabaseTestServiceImplementation(repository repository_database_test.DatabaseTestRepository) DatabaseTestService {
	return &DatabaseTestServiceImplementation{repository: repository}
}
