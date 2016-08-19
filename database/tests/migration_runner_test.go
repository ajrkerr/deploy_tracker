package Tests

import (
  Database "github.com/ajrkerr/deploy_tracker/database"
  . "github.com/ajrkerr/deploy_tracker/deploy_tracker"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

var _ = XDescribe("MigrationRunner", func() {
  var pgConfig = Database.PGConfig{
    Hostname: "localhost",
    Username: "vagrant",
    Database: "goTest",
    Port: 5434,
  }

  var connection = Database.PGConnect(pgConfig);

  It("Gets the current schema version", func() {
    var runner = Database.BuildMigrationRunner(connection)
    Expect(runner.SchemaVersion()).To(Equal(0))
  })

  Describe("Count", func() {
    It("Returns 0 when empty", func() {
      var repo = SqlDeployTrackerRepository(connection)

      Expect(repo.Count()).To(Equal(0))
    })
  })
})

func TestMigrationRunner(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Migration Runner Tests")
}
