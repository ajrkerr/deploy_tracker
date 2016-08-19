package DeployTracker

import (
  "time"
  "github.com/satori/go.uuid"
)

type deployInfo struct {
  deployTime  time.Time
  sha         string
  build       string
  environment string
  application string
  uuid        uuid.UUID
}

type DeployInfo interface {
  DeployTime() time.Time
  SHA() string
  Build() string
  Environment() string
  Application() string
  UUID() uuid.UUID
}

func (info deployInfo) DeployTime() time.Time {
  return info.deployTime;
}
func (info deployInfo) SHA() string {
  return info.sha;
}
func (info deployInfo) Build() string {
  return info.build;
}
func (info deployInfo) Environment() string {
  return info.environment;
}
func (info deployInfo) Application() string {
  return info.application;
}
func (info deployInfo) UUID() uuid.UUID {
  return info.uuid;
}