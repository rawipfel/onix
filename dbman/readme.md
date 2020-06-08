# DbMan - Onix Database Manager

DbMan is a command line interface application written in go, which aims at making it easy to track database versions, deploy schemas and functions and perform rolling upgrades.

DbMan enables development teams to keep their database scripts in a git repository and by the use of release manifests, automate the orchestration repetitive database release logic.

DbMan can be run in a docker container so that it can enable modern application deployment scenarios from a container platform such as Kubernetes.

Other functions such as backup/restore are currently being considered.

It currently supports PostgreSQL database but other database types are planned in the future using a plugin architecture.

## Command hierarchy

The following table presents a view of available cli commands:

| level 1 | level 2 | description | example |
|---|---|---|---|
| config | - | manages configuration sets | `dbman config [command]`|
| config | show | shows the content of the current configuration set | `dbman config show` |
| config | use | use a named configuration set. thsi command creates a new set if it does not exists and switches to it. | `dbman config use -n dev` |
| config | list | list all the available configuration sets indicating which is the selected one | `dbman config list` |
| config | set | sets a configuration value in the current set | `dbman config set schema.uri https://github.com/myrepo`|
| release | - | shows release information | `dbman release [command]` |
| release | plan | shows the release plan in the scripts repository | `dbman release plan` |
| release | info | show a specific release information | `dbman release info 0.0.4` |
| db | - | database maintenance tasks | `dbman db [command]` |
| db | init | initialises the database using the init manifest in the /init folder in the scripts repo | `dbman db init` |
| db | deploy | deploys the schema and objects for a particular release from the scripts repo | `dbman db deploy 0.0.4` |
| db | upgrade | upgrades the schema and objects to a particular release | `dbman db upgrade 0.0.4` |
| db | version | shows the version history in the tracking table | `dbman db version` |
| db | backup | takes a database backup | not currently supported |
| db | restore | restores a database backup | not currently supported |
| serve | - | starts dbman as an http service | `dbman serve` |

## Container Image Configuration

The container image can be configured by passing the following environment variables:

| var | description | default |
|---|---|---|
| `OX_DBM_HTTP_METRICS` | Whether prometheus `/metrics` endpoint is enabled. Only available if running dbman as an http service. | `true` |
| `OX_DBM_HTTP_AUTHMODE` | The authentication mode used by dbman http service. Acceptable values are `none` or `basic` for basic user authentication tokens. <br>Only available if running dbman as an http service. | `basic` |
| `OX_DBM_HTTP_PORT` | The port the http server is listening on.<br>Only available if running dbman as an http service. | `8085` |
| `OX_DBM_HTTP_USERNAME` | The username for the http service basic user authentication. <br>Only available if running dbman as an http service. | `admin` |
| `OX_DBM_HTTP_PASSWORD` | The password for the http service basic user authentication. <br>Only available if running dbman as an http service. | `0n1x` |
| `OX_DBM_DB_PROVIDER` | The database provider to use. Currently the only supported provider is PostgreSQL. | `pgsql` |
| `OX_DBM_DB_NAME` | The name of the database to manage. | `onix` |
| `OX_DBM_DB_HOST` | The database host | `localhost` |
| `OX_DBM_DB_PORT` | The database port | `5432` |
| `OX_DBM_DB_USERNAME` | The database user name | `onix` |
| `OX_DBM_DB_PASSWORD` | The database user password | `onix` |
| `OX_DBM_DB_ADMINUSER` | The database admin user | `postgres` |
| `OX_DBM_DB_ADMINPWD` | The database admin password | `onix` |
| `OX_DBM_SCHEMA_URI` | The root path of the database scripts. | `https://raw.githubusercontent.com/gatblau/ox-db/master` |
| `OX_DBM_SCHEMA_USERNAME` | The username for the scripts repository. | not supported |
| `OX_DBM_SCHEMA_TOKEN` | The token/password for the scripts repository. | not supported |

## Sample database scripts repository

For an example of the structure of the scripts repository required by dbman, [see here](https://github.com/gatblau/ox-db/).

## Container Registry

DbMan images can be found below:

| type | repo | image | tag |
|---|---|---|---|
| snapshot | docker.io | [gatblau/dbman-snapshot](https://hub.docker.com/repository/docker/gatblau/dbman-snapshot/general) |  latest |

**Note** DbMan is work in progress, it will replace and augment the logic in the Onix Web API
and allow a Kubernetes operator to orchestrate the required database admin tasks