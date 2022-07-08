# Stock API

[![Go](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-and-test.yml)

A Stock API is a REST API written in Golang where products (with stock) can be stored, retrieved, modified and deleted.

Note: _This API has been configured for `development` environment. To use in a `production` environment, further setup will be required._

## Install

```
git clone git@github.com:lucasvmiguel/stock-api.git
```

## Running the app

**Requirements:**

- [Golang](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)

1. Open a terminal and run the following command to start the persistence (database) required:

```bash
$ make persistence-up
```

2. In another terminal, start the application with the following command:

```bash
$ make run
```

## Testing

### Unit test

```
make test-unit
```

## Architecture

### Schema

![schema](/docs/schema.png)

### System Design

![system design](/docs/system-design.png)
![layers](/docs/layers.png)

### Folder/File struct

- `/cmd`: Main applications for this project.
- `/internal`: Private application and library code.
- `/internal/product`: Product domain, where every code related to product should be placed. (Inspired by [DDD](https://en.wikipedia.org/wiki/Domain-driven_design))
- `/pkg`: Library code that's ok to use by external applications (e.g., /pkg/mypubliclib).
- `/.github`: CI/CD from Github.
- `docker-compose.yml`: Used to spin up the persistence layer in development.
- `.env`: configures project.
- `Makefile`: Project's executable tasks.

Note: _inspired by https://github.com/golang-standards/project-layout_

### Stack

- Language: `Golang`
- API/REST framework: `chi`
- Database ORM: `Gorm`
- Config reader: `godotenv`
- Database: Postgres

## API Docs

In this section is described the REST API's endpoints (URL, request, response, etc).

<details>
<summary>Create product</summary>

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

</details>

<details>
<summary>Get all products</summary>

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

</details>

<details>
<summary>Get product by id</summary>

Endpoint to get a product by id

#### Request

```
Endpoint: [GET] /products/{id}

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

</details>

<details>
<summary>Update product by id</summary>

Endpoint that updates a product by id

#### Request

```
Endpoint: [PUT] /products/{id}

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

</details>

<details>
<summary>Delete product by id</summary>

Endpoint to delete a product by id

#### Request

```
Endpoint: [DELETE] /products/{id}

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

</details>

<br />

## Configuration

A file called `.env` has all config used in the project.

## CI/CD

The project uses Github CI to run tests, builds (and possibly deployments). You can see the badge below:
<br />
[![Go](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-and-test.yml)

Steps:

1. Set up Go
2. Build
3. Test
4. Log in to the Container registry (Github)
5. Build and push Docker images

## Important notes

- command `make docker-run` in `development` will only work correctly if the container's network is configured right. (Check more info [here](https://docs.docker.com/config/containers/container-networking/))

## Roadmap

- Remove `AutoMigrate` to implement some sort of manual migration system.
- Implement E2E tests.
- API docs are being described on the Readme. However, [OpenAPI](https://swagger.io/specification/) might be a good improvement in the future.
