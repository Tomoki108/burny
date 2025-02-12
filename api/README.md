# Burny API

REST API Server for Burny.

API DOC: http://localhost:1323/swagger/index.html

## teck stack

| category | TH                               |
| -------- | -------------------------------- |
| **FW**   | https://github.com/labstack/echo |
| **ORM**  | https://github.com/go-gorm/gorm  |
| **DB**   | PostgreSQL                       |

## Prerequisites

- Install [goenv](https://github.com/go-nv/goenv) and specified go version.

  ```shell
  brew install goenv
  goenv install 1.23.4
  ```

- Create `.envrc` from `.envrc.sample`. (Don't forget `direnv allow` after adding some env var.)

## How to run

```shell
docker compose up -d # for postgres container
go run .
```
