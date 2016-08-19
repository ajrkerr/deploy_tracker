package Tests

import (
  Database "github.com/ajrkerr/deploy_tracker/database"
  . "github.com/ajrkerr/deploy_tracker/deploy_tracker"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "testing"
)

var _ = XDescribe("DeployTrackerRepository", func() {
  var pgConfig = Database.PGConfig{
    Hostname: "localhost",
    Username: "vagrant",
    Database: "goTest",
    Port: 5434,
  }

  var connection = Database.PGConnect(pgConfig);

  It("Saves DeployInfo", func() {
    var repo = SqlDeployTrackerRepository(connection)

    var deployInfo = BuildDeployInfo("App", "Env", "Build", "SHA")

    repo.Persist(deployInfo)
    Expect(repo.Count()).To(Equal(1))
  })

  Describe("Count", func() {
    It("Returns 0 when empty", func() {
      var repo = SqlDeployTrackerRepository(connection)

      Expect(repo.Count()).To(Equal(0))
    })
  })
})

func TestRepository(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Repository Test")
}
