# Stock API

A Stock API is a REST API written in Golang where products (with stock) can be stored, retrieved, modified and deleted.

## Install

```
git clone git@github.com:lucasvmiguel/stock-api.git
```

_You must have Golang installed and configured to work with this API._

## Running the app

**Requirements:**

- [Golang](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)

1. Open a terminal and run the following command to start all persistence (database and queue) required:

```bash
$ make persistence-up
```

2. In another terminal, start the application with the following command:

```bash
$ make run
```

## Testing

### How to run unit

```
make test-unit
```

## Architecture

### Schema

![schema](/docs/schema.png)

### System Design

![system design](/docs/system-design.png)

### Folder/File struct

- `/cmd`: Main applications for this project.
- `/internal`: Part of the code that is not shareable (nor relevant) to other projects.
- `/internal/product`: Product domain, where every code related to product should be placed. (Inspired by [DDD](https://en.wikipedia.org/wiki/Domain-driven_design))
- `/pkg`: Part of the code that can be shared and used by other projects.
- `/.github`: CI/CD from Github.
- `docker-compose.yml`: Used to spin up the persistence layer in development and testing.
- `.env`: configures project.
- `Makefile`: Project's executable tasks.

### Stack

- Language: `Golang`
- API/REST framework: `chi`
- Database ORM: `Gorm`
- Config reader: `godotenv`

## API Docs

In this section is described the REST API's endpoints (URL, request, response, etc).

Note: _API docs are being described on the Readme. However, [OpenAPI](https://swagger.io/specification/) might be a good improvement in the future._

### Create product

Endpoint that creates a product

#### Request

```
Endpoint: [POST] /products

Headers:
  Content-Type: application/json

Body:
  {
    "Name": "Product name",
    "StockQuantity": 10
  }
```

#### Response

**Success**

```
Status: 201

Body:
  {
    "ID": 1,
    "CreatedAt": "2022-07-08T18:53:57.936433+01:00",
    "UpdatedAt": "2022-07-08T18:53:57.936433+01:00",
    "DeletedAt": null,
    "Name": "Product name",
    "Code": "70a17d32-a670-4396-9706-bd0940152fc7",
    "StockQuantity": 10
  }
```

**Bad Request**

```
Status: 400
```

**Internal Server Error**

```
Status: 500
```

### Get all products

Endpoint to get all products

#### Request

```
Endpoint: [GET] /products

Headers:
  Content-Type: application/json
```

#### Response

**Success**

```
Status: 200

Body:
  [
    {
      "ID": 1,
      "CreatedAt": "2022-07-08T18:53:57.936433+01:00",
      "UpdatedAt": "2022-07-08T18:53:57.936433+01:00",
      "DeletedAt": null,
      "Name": "foo",
      "Code": "70a17d32-a670-4396-9706-bd0940152fc7",
      "StockQuantity": 1
    }
  ]
```

**Internal Server Error**

```
Status: 500
```

### Get product by id

Endpoint to get a product by id

#### Request

```
Endpoint: [GET] /products/1

Headers:
  Content-Type: application/json
```

#### Response

**Success**

```
Status: 200

Body:
  {
    "ID": 1,
    "CreatedAt": "2022-07-08T18:53:57.936433+01:00",
    "UpdatedAt": "2022-07-08T18:53:57.936433+01:00",
    "DeletedAt": null,
    "Name": "foo",
    "Code": "70a17d32-a670-4396-9706-bd0940152fc7",
    "StockQuantity": 1
  }
```

**Not Found**

```
Status: 404
```

**Internal Server Error**

```
Status: 500
```

### Update product by id

Endpoint that updates a product by id

#### Request

```
Endpoint: [POST] /products

Headers:
  Content-Type: application/json

Body:
  {
    "Name": "new product name",
    "StockQuantity": 5
  }
```

#### Response

**Success**

```
Status: 200

Body:
  {
    "ID": 1,
    "CreatedAt": "2022-07-08T18:53:57.936433+01:00",
    "UpdatedAt": "2022-07-08T18:53:57.936433+01:00",
    "DeletedAt": null,
    "Name": "new product name",
    "Code": "70a17d32-a670-4396-9706-bd0940152fc7",
    "StockQuantity": 5
  }
```

**Bad Request**

```
Status: 400
```

**Not Found**

```
Status: 404
```

**Internal Server Error**

```
Status: 500
```

### Delete product by id

Endpoint to delete a product by id

#### Request

```
Endpoint: [DELETE] /products/1

Headers:
  Content-Type: application/json
```

#### Response

**Success**

```
Status: 204
```

**Not Found**

```
Status: 404
```

**Internal Server Error**

```
Status: 500
```

## Configuration

A file called `.env` has all config used in the project.

## CI/CD

The project uses Github CI to run tests, builds (and possibly deployments). You can see the badge below:
[![Node.js CI](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-an-test.yml/badge.svg)](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-and-test.yml)

Steps:

1. Set up Go
2. Build
3. Test
4. Log in to the Container registry (Github)
5. Build and push Docker images

## Roadmap

- Remove `AutoMigrate` to implement some sort of manual migration system.
