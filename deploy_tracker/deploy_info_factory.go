package DeployTracker

import (
  "time"
  "github.com/satori/go.uuid"
)

type deployInfoFactory struct {
  NowGenerator  func() time.Time
  UUIDGenerator func() uuid.UUID
}

type DeployInfoFactory interface {
  Build() DeployInfo
}

func BuildDeployInfoFactory(nowGenerator func() time.Time, uuidGenerator func() uuid.UUID) deployInfoFactory {
  return deployInfoFactory{nowGenerator, uuidGenerator}
}

func BuildDeployInfo(application string, environment string, build string, sha string) DeployInfo {
  var factory = BuildDeployInfoFactory(time.Now, uuid.NewV4)
  return factory.Build(application, environment, build, sha)
}

func (factory deployInfoFactory) Build(application string, environment string, build string, sha string) DeployInfo {
  return deployInfo{
    deployTime: factory.NowGenerator(),
    uuid: factory.UUIDGenerator(),
    application: application,
    environment: environment,
    build: build,
    sha: sha,
  }
}
