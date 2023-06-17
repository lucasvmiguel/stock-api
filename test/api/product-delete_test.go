package api

import (
	"net/http"
	"testing"

	"github.com/lucasvmiguel/integration"
	"github.com/lucasvmiguel/integration/assertion"
	"github.com/lucasvmiguel/integration/call"
	"github.com/lucasvmiguel/integration/expect"
)

func TestDeleteProduct_Successfully(t *testing.T) {
	reload()

	err := integration.Test(&integration.HTTPTestCase{
		Description: "TestGetAllProduct_Successfully",
		Request: call.Request{
			URL:    "http://localhost:8080/api/v1/products/1",
			Method: http.MethodDelete,
		},
		Response: expect.Response{
			StatusCode: http.StatusNoContent,
		},
		Assertions: []assertion.Assertion{
			&assertion.SQL{
				DB: DB,
				Query: call.Query{
					Statement: `
					SELECT name, stock_quantity, code FROM products WHERE deleted_at IS NULL
					`,
				},
				Result: expect.Result{
					{"name": "playstation 5", "stock_quantity": 2, "code": "b0553885-7d5b-4c9d-9ada-000000000002"},
					{"name": "nintendo switch", "stock_quantity": 3, "code": "b0553885-7d5b-4c9d-9ada-000000000003"},
					{"name": "xbox series s", "stock_quantity": 4, "code": "b0553885-7d5b-4c9d-9ada-000000000004"},
					{"name": "steam deck", "stock_quantity": 5, "code": "b0553885-7d5b-4c9d-9ada-000000000005"},
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
