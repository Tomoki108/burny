# Burny API

REST API Server for Burny.

API DOC: http://localhost:1323/swagger/index.html

## Teck Stack

| Category         | Tool                                       |
| ---------------- | ------------------------------------------ |
| **FW**           | https://github.com/labstack/echo           |
| **ORM**          | https://github.com/go-gorm/gorm            |
| **DB**           | PostgreSQL                                 |
| **validation**   | https://github.com/go-playground/validator |
| **di container** | https://github.com/uber-go/dig             |
| **e2e test**     | https://github.com/sebdah/goldie           |
| **doc**          | https://github.com/swaggo/swag             |

## Prerequisites

- Install [goenv](https://github.com/go-nv/goenv), specified go version and tools.

  ```shell
  brew install goenv
  goenv install 1.23.4
  go install github.com/swaggo/swag/cmd/swag@latest

  # deployment tools
  brew install skaffold
  brew install ko
  ```

## How to run

```shell
docker compose up -d # for postgres container
go run .
```

### Commands

```shell
# update API DOC
swag init

# push application image
skaffold build
# push application image && deploy to cloud run
skaffold run
```
