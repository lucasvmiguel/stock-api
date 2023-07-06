package api

import (
	"net/http"
	"testing"

	"github.com/lucasvmiguel/integration"
	"github.com/lucasvmiguel/integration/assertion"
	"github.com/lucasvmiguel/integration/call"
	"github.com/lucasvmiguel/integration/expect"
)

func TestUpdateProduct_Successfully(t *testing.T) {
	reload()

	err := integration.Test(&integration.HTTPTestCase{
		Description: "TestUpdateProduct_Successfully",
		Request: call.Request{
			URL:    "http://localhost:8080/api/v1/products/1",
			Method: http.MethodPut,
			Body: `{
				"name": "nintendo 64",
				"stock_quantity": 10
			}`,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `{
				"id": 1,
				"name":"nintendo 64",
				"stock_quantity":10,
				"code": "b0553885-7d5b-4c9d-9ada-000000000001",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			}`,
		},
		Assertions: []assertion.Assertion{
			&assertion.SQL{
				DB: DB,
				Query: call.Query{
					Statement: `
					SELECT name, stock_quantity FROM products ORDER BY id LIMIT 1
					`,
				},
				Result: expect.Result{
					{"name": "nintendo 64", "stock_quantity": 10},
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
