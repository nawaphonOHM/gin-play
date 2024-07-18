package repository_database_test

import "github.com/jmoiron/sqlx"

type DatabaseTestRepositoryImplementation struct {
	databaseWrapper *sqlx.DB
}

func (r *DatabaseTestRepositoryImplementation) GetData() *Response {
	//TODO implement me
	panic("implement me")
}

func NewDatabaseTestRepositoryImplementation(db *sqlx.DB) DatabaseTestRepository {
	return &DatabaseTestRepositoryImplementation{databaseWrapper: db}
}
