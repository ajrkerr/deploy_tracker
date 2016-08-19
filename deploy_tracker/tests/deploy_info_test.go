package Tests

import (
  . "github.com/ajrkerr/deploy_tracker/deploy_tracker"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "time"
  "testing"
  "github.com/satori/go.uuid"
)

var _ = Describe("BuidDeployInfo", func() {
  It("Builds deploy infos", func() {
    var deployInfo = BuildDeployInfo("App", "Environment", "Build", "SHA");

    Expect(deployInfo.UUID()).ToNot(Equal(uuid.Nil))
    Expect(deployInfo.DeployTime().IsZero()).To(BeFalse())
    Expect(deployInfo.Application()).To(Equal("App"))
    Expect(deployInfo.Build()).To(Equal("Build"))
    Expect(deployInfo.SHA()).To(Equal("SHA"))
    Expect(deployInfo.Environment()).To(Equal("Environment"))
  })
});

var _ = Describe("BuildDeployInfoFactory", func() {
  It("Builds DeployInfo with the correct parameters", func() {
    var testNow = time.Date(2015, 1, 1, 12, 0, 0, 0, time.UTC)
    var testUuid, _ = uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")

    var getNow = func () time.Time { return testNow }
    var getUUID = func () uuid.UUID { return testUuid }

    var factory = BuildDeployInfoFactory(getNow, getUUID)
    var deployInfo = factory.Build("App", "Environment", "Build", "SHA")

    Expect(deployInfo.DeployTime()).To(Equal(testNow))
    Expect(deployInfo.Application()).To(Equal("App"))
    Expect(deployInfo.Build()).To(Equal("Build"))
    Expect(deployInfo.SHA()).To(Equal("SHA"))
    Expect(deployInfo.Environment()).To(Equal("Environment"))
  })

  It("has a GUID", func() {
    var factory = BuildDeployInfoFactory(time.Now, uuid.NewV4)
    var deployInfo = factory.Build("App", "Environment", "Build", "SHA")

    Expect(deployInfo.UUID()).ToNot(Equal(uuid.Nil))
  })

  It("generates a time using the real clock", func() {
    var factory = BuildDeployInfoFactory(time.Now, uuid.NewV4)
    var deployInfo = factory.Build("App", "Environment", "Build", "SHA")

    Expect(deployInfo.DeployTime().IsZero()).To(BeFalse())
  })
})

func TestDeployTracker(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Deploy Tracker Tests")
}
