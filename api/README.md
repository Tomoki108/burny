# Burny API

REST API Server for Burny.

API DOC: https://dev-api.burny.page/swagger/index.html

## Teck Stack

| Category          | Tool                                       |
| ----------------- | ------------------------------------------ |
| **FW**            | https://github.com/labstack/echo           |
| **ORM**           | https://github.com/go-gorm/gorm            |
| **DB**            | PostgreSQL                                 |
| **Architecture**  | Clean Architecture                         |
| **Validation**    | https://github.com/go-playground/validator |
| **DI Container**  | https://github.com/uber-go/dig             |
| **Event Bus**     | https://github.com/asaskevich/EventBus     |
| **Scenario Test** | https://github.com/sebdah/goldie           |
| **Swagger**       | https://github.com/swaggo/swag             |

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
```

### Other Commands

```shell
# update API DOC
swag init

# run scenario tests
go test ./scenario
# update golden files of scenario tests
go test ./scenario -update
```
