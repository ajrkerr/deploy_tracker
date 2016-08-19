package Database

type MigrationRunner interface {
  SchemaVersion() int
}

type migrationRunner struct {
  connection SqlConnection
}

func BuildMigrationRunner(connection SqlConnection) MigrationRunner {
  return migrationRunner{connection: connection}
}

func (runner migrationRunner) SchemaVersion() int {
  return -1
}