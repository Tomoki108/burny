# Burny API

REST API Server for Burny.

API DOC: https://api.burny.page/swagger/index.html

## Teck Stack

| Category          | Tool                                                                  |
| ----------------- | --------------------------------------------------------------------- |
| **FW**            | [labstack/echo](https://github.com/labstack/echo)                     |
| **ORM**           | [go-gorm/gorm](https://github.com/go-gorm/gorm)                       |
| **DB**            | PostgreSQL                                                            |
| **Architecture**  | Clean Architecture                                                    |
| **Validation**    | [go-playground/validator](https://github.com/go-playground/validator) |
| **DI Container**  | [uber-go/dig](https://github.com/uber-go/dig)                         |
| **Event Bus**     | [asaskevich/EventBus](https://github.com/asaskevich/EventBus)         |
| **Scenario Test** | [sebdah/goldie](https://github.com/sebdah/goldie)                     |
| **Swagger**       | [swaggo/swag](https://github.com/swaggo/swag)                         |

## How to run

```shell
docker compose up -d # for postgres container
go run .
```

### Prerequisites

- Install [goenv](https://github.com/go-nv/goenv), specified go version and tools.

```shell
brew install goenv
goenv install 1.23.4
go install github.com/swaggo/swag/cmd/swag@latest
```

- Install and setting [direnv](https://github.com/direnv/direnv). Then create envrc.

```shell
cp .envrc.sample .envrc # AWS access keys must be set
direnv allow
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
