# opensirene [![Build Status](https://travis-ci.org/jclebreton/opensirene.svg?branch=v2)](https://travis-ci.org/jclebreton/opensirene) [![codecov](https://codecov.io/gh/jclebreton/opensirene/branch/master/graph/badge.svg)](https://codecov.io/gh/jclebreton/opensirene)
French company database based on French government open data

## Getting Started

### Build
```
$ dep ensure
$ go run main.go
```

## Configuration

This project can be configured using both a yaml configuration file and
environment variables (for most of the configuration fields). Environment
variables have a higher priority than the configuration file, which means you
can override almost any value of the configuration file using them. 

yml
---
```
logger:
  level: debug
  format: text

server:
  host: 127.0.0.1
  port: 8080
  debug: false
  cors:
    permissive_mode: true
    enabled: true

database:
  user: xx
  password: xx
  name: opensirene
  host: 127.0.0.1
  port: 5432

prometheus:
  prefix: opensirene

crontab:
  download_path: "downloads"
  every_x_hours: 3

```


| Field                       | Type     | Description                                               | Environment Variable | Default        | Example        |
|-----------------------------|----------|-----------------------------------------------------------|----------------------|----------------|----------------|
| logger.level                | string   | Global log level                                          | `LOGLEVEL`           | "info"         | "debug"        |
| logger.format               | string   | Log format (text, json)                                   | `LOGFORMAT`          | "text"         | "json"         |
| server.host                 | string   | Host on which the server will listen                      | `SERVER_HOST`        | "127.0.0.1"    | "127.0.0.1"    |
| server.port                 | int      | Port on which the server will listen                      | `SERVER_PORT`        | 8080           | 8080           |
| server.debug                | bool     | Debug mode                                                | `SERVER_DEBUG`       | false          | true           |
| server.cors.allow_origins   | []string | Array of accepted origins                                 | -                    | -              | -              |
| server.cors.permissive_mode | bool     | Accept every origin and overrides the allow_origins field | `CORS_PERMISSIVE`    | false          | true           |
| database.user               | string   | User used to connect to the DB                            | `DB_USER`            | "sir"          | "sir"          |
| database.password           | string   | Password associated to the user                           | `DB_PASSWORD`        | -              | -              |
| database.host               | string   | Host on which the DB listens                              | `DB_HOST`            | "127.0.0.1"    | "127.0.0.1"    |
| database.port               | int      | Port on which the DB listens                              | `DB_PORT`            | 5432           | 5432           |
| database.name               | string   | Database name to use                                      | `DB_NAME`            | "opensirenedb" | "opensirenedb" |
| database.sslmode            | string   | Use the SSL mode                                          | `DB_SSL_MODE`        | "disable"      | "disable"      |
| prometheus.prefix           | string   | Prefix the prometheus metrics                             | `PROMETHEUS_PREFIX`  | "opensirene"   | "opensirene"   |
| crontab.download_path       | string   | Downloads path                                            | `DOWNLOAD_PATH`      | "downloads"    | "/tmp"         |
| crontab.every_x_hours       | uint64   | Crontab interval (in hours)                               | `EVERY_X_HOURS`      | 3              | 1              |

## Steps to finish (redoc)

1. Enable [Travis](https://docs.travis-ci.com/user/getting-started/#To-get-started-with-Travis-CI%3A) for your repository (**note**: you already have `.travis.yml` file)
2. [Create GitHub access token](https://help.github.com/articles/creating-an-access-token-for-command-line-use/); check `public_repo` on `Select scopes` section.
3. Use the token value as a value for [Travis environment variable](https://docs.travis-ci.com/user/environment-variables/#Defining-Variables-in-Repository-Settings) with the name `GH_TOKEN`
4. Make a test commit to trigger CI: `git commit --allow-empty -m "Test Travis CI" && git push`
5. Wait until Travis build is finished. You can check progress by clicking on the `Build Status` badge at the top
6. If you did everything correct, https://samsaggace.github.io/opensirene/ will lead to your new docs
7. **[Optional]** You can setup [custom domain](https://help.github.com/articles/using-a-custom-domain-with-github-pages/) (just create `web/CNAME` file)
8. Start writing/editing your OpenAPI spec: check out [usage](#usage) section below
9. **[Optional]** If you document public API consider adding it into [APIs.guru](https://APIs.guru) directory using [this form](https://apis.guru/add-api/).
10. Delete this section :smile:

## Links

- Documentation(ReDoc): https://samsaggace.github.io/opensirene/
- Look full spec:
    + JSON https://samsaggace.github.io/opensirene/swagger.json
    + YAML https://samsaggace.github.io/opensirene/swagger.yaml
- Preview spec version for branch `[branch]`: https://samsaggace.github.io/opensirene/preview/[branch]

**Warning:** All above links are updated only after Travis CI finishes deployment

## Working on specification
### Install

1. Install [Node JS](https://nodejs.org/)
2. Clone repo and `cd`
    + Run `npm install`

### Usage

1. Run `npm start`
2. Checkout console output to see where local server is started. You can use all [links](#links) (except `preview`) by replacing https://samsaggace.github.io/opensirene/ with url from the message: `Server started <url>`
3. Make changes using your favorite editor or `swagger-editor` (look for URL in console output)
4. All changes are immediately propagated to your local server, moreover all documentation pages will be automagically refreshed in a browser after each change
**TIP:** you can open `swagger-editor`, documentation and `swagger-ui` in parallel
5. Once you finish with the changes you can run tests using: `npm test`
6. Share you changes with the rest of the world by pushing to GitHub :smile:
