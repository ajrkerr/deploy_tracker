package Database

import "database/sql"

type SqlConnection interface {
  Query(query string, args ...interface{}) (*sql.Rows, error)

  CreateDatabase(dbName string)
  DatabaseExists(dbName string) bool
}