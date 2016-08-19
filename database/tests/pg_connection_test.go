package Tests

import (
  Database "github.com/ajrkerr/deploy_tracker/database"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

var _ = Describe("PgConnection", func() {
  var pgConfig = Database.PGConfig{
    Hostname: "localhost",
    Username: "vagrant",
    Database: "goTest",
    Port: 5434,
  }

  var connection = Database.PGConnect(pgConfig);
  var dbName = "TestDB"

  BeforeEach(func() {
    connection.DropDatabase(dbName)
  })

  It("Creates a database", func() {
    connection.CreateDatabase(dbName)

    Expect(connection.DatabaseExists(dbName)).To(BeTrue())
  })

  It("Drops databases", func() {
    connection.CreateDatabase(dbName)
    connection.DropDatabase(dbName)

    Expect(connection.DatabaseExists(dbName)).To(BeFalse())
  })
})

func TestPGConnection(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "PG Connection Tests")
}
