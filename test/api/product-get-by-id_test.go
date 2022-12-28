package api

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/lucasvmiguel/integration"
	"github.com/lucasvmiguel/integration/call"
	"github.com/lucasvmiguel/integration/expect"
)

func TestGetByIDProduct_Successfully(t *testing.T) {
	reload()

	var id int
	row := DB.QueryRow("SELECT id FROM products LIMIT 1")
	row.Scan(&id)

	err := integration.Test(integration.TestCase{
		Description: "TestGetAllProduct_Successfully",
		Request: call.Request{
			URL:    fmt.Sprintf("http://localhost:8080/products/%d", id),
			Method: http.MethodGet,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `{
				"name":"foo",
				"stock_quantity":1
			}`,
			IgnoreBodyFields: []string{
				"id", "code", "created_at", "updated_at",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}