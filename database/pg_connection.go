package Database

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
  "log"
)

type PGConfig struct {
  Hostname string
  Username string
  Password string
  Database string
  Port     int
}

func (config PGConfig) toConnectString() string {
  return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=postgres", config.Username, config.Password, config.Hostname, config.Port)
}

func PGConnect(config PGConfig) pgConnection {
  var connectionString = config.toConnectString()
  var db, err = sql.Open("postgres", connectionString)

  if err != nil {
    log.Fatal(err)
  }

  return pgConnection{db: db}
}

type pgConnection struct {
  db *sql.DB
}

func (conn pgConnection) Query(query string, args ...interface{}) (*sql.Rows, error) {
  var rows, err = conn.db.Query(query, args...)

  if err != nil {
    log.Print(err)
  }

  return rows, err
}

func (conn pgConnection) CreateDatabase(dbName string) {
  conn.Query(fmt.Sprintf("CREATE DATABASE \"%s\"", dbName))
}

func (conn pgConnection) DropDatabase(dbName string) {
  conn.Query(fmt.Sprintf("DROP DATABASE IF EXISTS \"%s\"", dbName))
}

func (conn pgConnection) DatabaseExists(dbName string) bool {
  var rows, err = conn.Query(fmt.Sprintf("SELECT COUNT(*) AS count FROM pg_database WHERE datname = '%s'", dbName))
  defer rows.Close()

  if err != nil { return false }

  if rows.Next() {
    var count int
    rows.Scan(&count)
    return count > 0
  }

  return false
}