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

func TestDeleteProduct_Successfully(t *testing.T) {
	DB.Exec("DELETE FROM products")
	DB.Exec("INSERT INTO products (name, stock_quantity) VALUES ('bar', 2);")

	var id int
	row := DB.QueryRow("SELECT id FROM products LIMIT 1")
	row.Scan(&id)

	err := integration.Test(integration.TestCase{
		Description: "TestGetAllProduct_Successfully",
		Request: call.Request{
			URL:    fmt.Sprintf("http://localhost:8080/products/%d", id),
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
					SELECT * FROM products WHERE deleted_at IS NULL
					`,
				},
				Result: expect.Result{},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
