package service_database_test

type DatabaseTestServiceImplementation struct {
}

func (s *DatabaseTestServiceImplementation) GetData() *Response {
	//TODO implement me
	panic("implement me")
}

func NewDatabaseTestServiceImplementation() DatabaseTestService {
	return &DatabaseTestServiceImplementation{}
}
