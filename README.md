# Stock API

[![Go](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-and-test.yml)

## Description

A Stock API is a REST API written in Go where products can be created, read, updated and deleted.

Note: _This API has been configured for `development` environment. To use in a `production` environment, further setup will be required._

## Installing the app

**Requirements:**

- [Go](https://go.dev/) (>= 1.20)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)

To install the go modules required to run the application, run the following command:

```bash
git clone git@github.com:lucasvmiguel/stock-api.git
cd stock-api
make install
```

## Running the app

1. Open a terminal and run the following command to start the persistence (database) required:

```bash
make persistence-up
```

2. In another terminal, start the application with the following command:

```bash
make run
```

## Testing

### Unit test

```
make test-unit
```

### Integration test

1. Open a terminal and run the following command to start the persistence (database) required:

```bash
make persistence-up
```

2. In another terminal, run the integration test with the following command:

```bash
make test-integration
```

### Stress test

**Requirements:**

- [Apache AB](https://httpd.apache.org/docs/2.4/programs/ab.html)

1. Open a terminal and run the following command to start the persistence (database) required:

```bash
make persistence-up
```

2. In a new terminal, start the application with the following command:

```bash
make run
```

3. In another terminal, run the stress test with the following command

```bash
make test-stress
```

## Configuration

- To configure how the app will run, check the following file: [.env](.env)
- To configure how the app will be built and released, check the following file: [Makefile](Makefile)

## Architecture

This section describes what are the goals of the system and some of its design and implementation.

### Requirements

The following list shows all user requirements implemented by the system.

- As a user, I can fetch all products using a REST API.
- As a user, I can fetch a paginated list of products using a REST API.
- As a user, I can fetch a product by its id using a REST API.
- As a user, I can create a product using a REST API.
- As a user, I can update some fields (`name` and/or `code`) of a product by its id using a REST API.
- As a user, I can delete a product by its id using a REST API.

### Schema

The following picture shows all the entities of the system.

![schema](/assets/schema.png)

### System Design

The following pictures shows some of the details of how the system is designed and implemented.

![system design](/assets/system-design.png)
![layers](/assets/layer.png)

### Folder/File structure

- `/cmd`: Main applications for this project.
- `/cmd/api`: Package responsible for starting the API application.
- `/cmd/api/starter`: Package containing all code required for starting the API application. (eg: config, routes, etc)
- `/internal`: Private application and library code.
- `/internal/product`: Product domain, where every code related to product should be placed. (Inspired by [DDD](https://en.wikipedia.org/wiki/Domain-driven_design))
- `/pkg`: Library code that's ok to use by external applications (eg: `/pkg/mypubliclib`).
- `/test`: Integration tests that run with external apps. (eg: database)
- `/.github`: CI/CD from Github.
- `docker-compose.yml`: Used to spin up the persistence layer in development and testing.
- `.env`: configures project.
- `Makefile`: Project's executable tasks.

Note: _inspired by https://github.com/golang-standards/project-layout_

### Stack

- Language: [Go](https://go.dev/)
- API/REST framework: [chi](https://github.com/go-chi/chi)
- Database ORM: [GORM](https://gorm.io/)
- Config reader: [godotenv](https://github.com/joho/godotenv)
- Database: [Postgres](https://www.postgresql.org/)
- Message Queue: [Asynq](https://github.com/hibiken/asynq)

## API Docs

In this section is described the REST API's endpoints (URL, request, response, etc).

<details>
<summary>Create product</summary>

Endpoint that creates a product

#### Request

```
Endpoint: [POST] /api/v1/products

Headers:
  Content-Type: application/json

Body:
  {
    "name": "Product name",
    "stock_quantity": 10
  }
```

#### Response

**Success**

```
Status: 201

Body:
  {
    "id": 1,
    "name": "Product name",
    "code": "70a17d32-a670-4396-9706-bd0940152fc7",
    "stock_quantity": 10,
    "created_at": "2022-07-08T18:53:57.936433+01:00",
    "updated_at": "2022-07-08T18:53:57.936433+01:00"
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
<summary>Get products paginated</summary>

Endpoint to get products paginated

#### Request

##### Query Parameters

- `cursor`: use the response's `next_cursor` field
- `limit`: limit of products to be returned (min=1, max=100)

```
Endpoint: [GET] /api/v1/products?limit=10&cursor=2

Headers:
  Content-Type: application/json
```

#### Response

**Success**

```
Status: 200

Body:
  {
    "items": [
      {
        "id": 1,
        "name": "foo",
        "code": "70a17d32-a670-4396-9706-bd0940152fc7",
        "stock_quantity": 1,
        "created_at": "2022-07-08T18:53:57.936433+01:00",
        "updated_at": "2022-07-08T18:53:57.936433+01:00"
      }
    ],
    "next_cursor": 2
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

Endpoint to get all products (does not have pagination)

#### Request

```
Endpoint: [GET] /api/v1/products/all

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
      "id": 1,
      "name": "foo",
      "code": "70a17d32-a670-4396-9706-bd0940152fc7",
      "stock_quantity": 1,
      "created_at": "2022-07-08T18:53:57.936433+01:00",
      "updated_at": "2022-07-08T18:53:57.936433+01:00"
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
Endpoint: [GET] /api/v1/products/{id}

Headers:
  Content-Type: application/json
```

#### Response

**Success**

```
Status: 200

Body:
  {
    "id": 1,
    "name": "foo",
    "code": "70a17d32-a670-4396-9706-bd0940152fc7",
    "stock_quantity": 1,
    "created_at": "2022-07-08T18:53:57.936433+01:00",
    "updated_at": "2022-07-08T18:53:57.936433+01:00"
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
Endpoint: [PUT] /api/v1/products/{id}

Headers:
  Content-Type: application/json

Body:
  {
    "name": "new product name",
    "stock_quantity": 5
  }
```

#### Response

**Success**

```
Status: 200

Body:
  {
    "id": 1,
    "name": "new product name",
    "code": "70a17d32-a670-4396-9706-bd0940152fc7",
    "stock_quantity": 5,
    "created_at": "2022-07-08T18:53:57.936433+01:00",
    "updated_at": "2022-07-08T18:53:57.936433+01:00"
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
Endpoint: [DELETE] /api/v1/products/{id}

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

## Configuration

A file called `.env` has all config used in the project.

In the future, a service like [Doppler](https://www.doppler.com/) or [Vault](https://www.vaultproject.io/) could (and should) be used in the project.

## CI/CD

The project uses Github CI to run tests, builds (and possibly deployments). You can see the badge below:
<br />
[![Go](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/lucasvmiguel/stock-api/actions/workflows/build-and-test.yml)

Steps:

1. Set up Go
2. Build
3. Unit Test
4. Integration Test
5. Log in to the Container registry (Github)
6. Build and push Docker images

## Important notes

- command `make docker-run` in `development` will only work correctly if the container's network is configured right. (More info [here](https://docs.docker.com/config/containers/container-networking/))

## Roadmap

- `Improvement`: If it's needed to add more entities (eg: [Product](internal/product/entity/product.go)), we might need to centralize all entities in just one package. (Something like a `entity` package) That way, we would prevent cycle dependencies. (Check [this link](https://www.reddit.com/r/golang/comments/vcy5xq/ddd_file_structure_cyclic_dependencies/))
- `Improvement`: API docs are being described on the Readme. However, [OpenAPI](https://swagger.io/specification/) might be a good improvement in the future.
- `Improvement`: Using a secret management service like [Doppler](https://www.doppler.com/) or [Vault](https://www.vaultproject.io/)
