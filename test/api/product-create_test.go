package api

import (
	"net/http"
	"testing"

	"github.com/lucasvmiguel/integration"
	"github.com/lucasvmiguel/integration/assertion"
	"github.com/lucasvmiguel/integration/call"
	"github.com/lucasvmiguel/integration/expect"
)

func TestCreateProduct_Successfully(t *testing.T) {
	reload()

	err := integration.Test(integration.TestCase{
		Description: "TestCreateProduct_Successfully",
		Request: call.Request{
			URL:    "http://localhost:8080/products",
			Method: http.MethodPost,
			Body: `{
				"name": "product test",
				"stock_quantity": 5
			}`,
		},
		Response: expect.Response{
			StatusCode: http.StatusCreated,
			Body: `{
				"id": "<<PRESENCE>>",
				"name":"product test",
				"stock_quantity":5,
				"code": "<<PRESENCE>>",
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>"
			}`,
		},
		Assertions: []assertion.Assertion{
			&assertion.SQL{
				DB: DB,
				Query: call.Query{
					Statement: `
					SELECT name, stock_quantity FROM products ORDER BY id DESC LIMIT 1
					`,
				},
				Result: expect.Result{
					{"name": "product test", "stock_quantity": 5},
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
