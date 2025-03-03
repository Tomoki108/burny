# Burny API

REST API Server for Burny.

API DOC: http://dev.burny.page/swagger/index.html

## Teck Stack

| Category         | Tool                                       |
| ---------------- | ------------------------------------------ |
| **FW**           | https://github.com/labstack/echo           |
| **ORM**          | https://github.com/go-gorm/gorm            |
| **DB**           | PostgreSQL                                 |
| **Validation**   | https://github.com/go-playground/validator |
| **DI Container** | https://github.com/uber-go/dig             |
| **E2E Test**     | https://github.com/sebdah/goldie           |
| **Swagger**      | https://github.com/swaggo/swag             |

## How to run

```shell
docker compose up -d # for postgres container
go run .
```

### Prerequisites

Install [goenv](https://github.com/go-nv/goenv), specified go version and tools.

```shell
brew install goenv
goenv install 1.23.4
go install github.com/swaggo/swag/cmd/swag@latest

# deployment tools
brew install skaffold
brew install ko
```

### Other Commands

```shell
# update API DOC
swag init

# push application image
skaffold build
# push application image && deploy to cloud run
skaffold run
```
