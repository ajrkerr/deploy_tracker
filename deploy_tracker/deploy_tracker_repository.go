package DeployTracker

import Database "github.com/ajrkerr/deploy_tracker/database"

type DeployTrackerRepository interface {
  Persist(info DeployInfo)
  Count() int
  Top10() []DeployInfo
}

func SqlDeployTrackerRepository(connection Database.SqlConnection) sqlDeployTrackerRepository {
  return sqlDeployTrackerRepository{connection: connection}
}

type sqlDeployTrackerRepository struct {
  connection Database.SqlConnection
}

func (repo sqlDeployTrackerRepository) Persist(info DeployInfo) {
  repo.connection.Query("")
}

func (repo sqlDeployTrackerRepository) Count() int {
  var rows, _ = repo.connection.Query("SELECT COUNT(*) AS count FROM deploy_infos");
  defer rows.Close()

  for rows.Next() {
    var count int
    rows.Scan(&count)

    return count
  }

  return -1
}

func (repo sqlDeployTrackerRepository) Top10() []DeployInfo {
  return []DeployInfo{
    BuildDeployInfo("App1", "Env1", "Build1", "SHA1"),
    BuildDeployInfo("App2", "Env2", "Build2", "SHA2"),
    BuildDeployInfo("App3", "Env3", "Build3", "SHA3"),
    BuildDeployInfo("App4", "Env4", "Build4", "SHA4"),
    BuildDeployInfo("App5", "Env5", "Build5", "SHA5"),
    BuildDeployInfo("App6", "Env6", "Build6", "SHA6"),
    BuildDeployInfo("App7", "Env7", "Build7", "SHA7"),
    BuildDeployInfo("App8", "Env8", "Build8", "SHA8"),
    BuildDeployInfo("App9", "Env9", "Build9", "SHA9"),
    BuildDeployInfo("App10", "Env10", "Build10", "SHA10"),
  }
}
