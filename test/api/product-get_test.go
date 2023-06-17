package api

import (
	"net/http"
	"testing"

	"github.com/lucasvmiguel/integration"
	"github.com/lucasvmiguel/integration/call"
	"github.com/lucasvmiguel/integration/expect"
)

func TestGet_Successfully(t *testing.T) {
	reloadWithMoreRows()

	err := integration.Test(&integration.HTTPTestCase{
		Description: "TestGet_Successfully page 1",
		Request: call.Request{
			URL:    "http://localhost:8080/api/v1/products?limit=2",
			Method: http.MethodGet,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `{
				"items": [
					{
						"id": 1,
						"name": "foo 1",
						"code": "b0553885-7d5b-4c9d-9ada-000000000001",
						"stock_quantity": 1,
						"created_at": "<<PRESENCE>>",
						"updated_at": "<<PRESENCE>>"
					},
					{
						"id": 2,
						"name": "foo 2",
						"code": "b0553885-7d5b-4c9d-9ada-000000000002",
						"stock_quantity": 2,
						"created_at": "<<PRESENCE>>",
						"updated_at": "<<PRESENCE>>"
					}
				],
				"next_cursor": 2
			}`,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	err = integration.Test(&integration.HTTPTestCase{
		Description: "TestGet_Successfully page 2",
		Request: call.Request{
			URL:    "http://localhost:8080/api/v1/products?limit=2&cursor=2",
			Method: http.MethodGet,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `{
				"items": [
					{
						"id": 3,
						"name": "foo 3",
						"code": "b0553885-7d5b-4c9d-9ada-000000000003",
						"stock_quantity": 3,
						"created_at": "<<PRESENCE>>",
						"updated_at": "<<PRESENCE>>"
					},
					{
						"id": 4,
						"name": "foo 4",
						"code": "b0553885-7d5b-4c9d-9ada-000000000004",
						"stock_quantity": 4,
						"created_at": "<<PRESENCE>>",
						"updated_at": "<<PRESENCE>>"
					}
				],
				"next_cursor": 4
			}`,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	err = integration.Test(&integration.HTTPTestCase{
		Description: "TestGet_Successfully page 3",
		Request: call.Request{
			URL:    "http://localhost:8080/api/v1/products?limit=2&cursor=4",
			Method: http.MethodGet,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `{
				"items": [
					{
						"id": 5,
						"name": "foo 5",
						"code": "b0553885-7d5b-4c9d-9ada-000000000005",
						"stock_quantity": 5,
						"created_at": "<<PRESENCE>>",
						"updated_at": "<<PRESENCE>>"
					}
				],
				"next_cursor": 5
			}`,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	err = integration.Test(&integration.HTTPTestCase{
		Description: "TestGet_Successfully page end",
		Request: call.Request{
			URL:    "http://localhost:8080/api/v1/products?limit=2&cursor=5",
			Method: http.MethodGet,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `{
				"items": [],
				"next_cursor": null
			}`,
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
