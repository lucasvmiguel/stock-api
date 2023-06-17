package api

import (
	"net/http"
	"testing"

	"github.com/lucasvmiguel/integration"
	"github.com/lucasvmiguel/integration/call"
	"github.com/lucasvmiguel/integration/expect"
)

func TestGetAllProduct_Successfully(t *testing.T) {
	reload()

	err := integration.Test(&integration.HTTPTestCase{
		Description: "TestGetAllProduct_Successfully",
		Request: call.Request{
			URL:    "http://localhost:8080/api/v1/products/all",
			Method: http.MethodGet,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `[
				{
					"id": 1,
					"name": "foo",
					"code": "b0553885-7d5b-4c9d-9ada-000000000001",
					"stock_quantity": 1,
					"created_at": "<<PRESENCE>>",
					"updated_at": "<<PRESENCE>>"
				},
				{
					"id": 2,
					"name": "bar",
					"code": "b0553885-7d5b-4c9d-9ada-000000000002",
					"stock_quantity": 2,
					"created_at": "<<PRESENCE>>",
					"updated_at": "<<PRESENCE>>"
				}
			]`,
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
