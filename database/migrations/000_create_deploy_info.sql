CREATE TABLE deploy_info (
    uuid varchar(32),
    deploy_time timestamp,
    sha varchar(64),
    application varchar(255),
    environment varchar(255),
    build varchar(255),

    PRIMARY KEY(uuid)
);

CREATE UNIQUE INDEX deploy_info_uuid_idx ON deploy_info (uuid);
CREATE UNIQUE INDEX deploy_info_deploy_time_idx ON deploy_info (deploy_time);
CREATE UNIQUE INDEX deploy_info_application_idx ON deploy_info (application);
CREATE UNIQUE INDEX deploy_info_environment_idx ON deploy_info (environment);