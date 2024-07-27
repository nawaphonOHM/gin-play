package repository_database_test

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type DatabaseTestRepositoryImplementation struct {
	databaseWrapper *sqlx.DB
}

func (r *DatabaseTestRepositoryImplementation) GetData() (*Response, error) {
	query := "SELECT * FROM \"User\""

	rows, err := r.databaseWrapper.Query(query)
	if err != nil {
		return nil
	}

	response := new(Response)

	for rows.Next() {
		row := make([]string, 0)
		err = rows.Scan(&row)
		if err != nil {
			log.Fatal(err)
		}
		response.Data = append(response.Data, row)
	}

	return response
}

func NewDatabaseTestRepositoryImplementation(db *sqlx.DB) DatabaseTestRepository {
	return &DatabaseTestRepositoryImplementation{databaseWrapper: db}
}
