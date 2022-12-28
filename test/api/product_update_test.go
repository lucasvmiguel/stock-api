package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/lucasvmiguel/integration"
	"github.com/lucasvmiguel/integration/assertion"
	"github.com/lucasvmiguel/integration/call"
	"github.com/lucasvmiguel/integration/expect"
)

func TestUpdateProduct_Successfully(t *testing.T) {
	DB.Exec("DELETE FROM products")
	DB.Exec("INSERT INTO products (name, stock_quantity) VALUES ('bar', 2);")

	var id int
	row := DB.QueryRow("SELECT id FROM products LIMIT 1")
	row.Scan(&id)

	err := integration.Test(integration.TestCase{
		Description: "TestCreateProduct_Successfully",
		Request: call.Request{
			URL:    fmt.Sprintf("http://localhost:8080/products/%d", id),
			Method: http.MethodPut,
			Body: `{
				"name": "foo",
				"stock_quantity": 10
			}`,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `{
				"name":"foo",
				"stock_quantity":10
			}`,
			IgnoreBodyFields: []string{"id", "code", "created_at", "updated_at"},
		},
		Assertions: []assertion.Assertion{
			&assertion.SQL{
				DB: DB,
				Query: call.Query{
					Statement: `
					SELECT name, stock_quantity FROM products
					`,
				},
				Result: expect.Result{
					{"name": "foo", "stock_quantity": 10},
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
