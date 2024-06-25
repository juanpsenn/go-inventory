package sqlite

import "database/sql"

// "./users.db"
func Connect(dbConnectionString string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbConnectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
