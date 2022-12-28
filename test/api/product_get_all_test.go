package api

import (
	"net/http"
	"testing"

	"github.com/lucasvmiguel/integration"
	"github.com/lucasvmiguel/integration/call"
	"github.com/lucasvmiguel/integration/expect"
)

func TestGetAllProduct_Successfully(t *testing.T) {
	DB.Exec("DELETE FROM products")
	DB.Exec("INSERT INTO products (name) VALUES ('foo');")
	DB.Exec("INSERT INTO products (name) VALUES ('bar');")

	err := integration.Test(integration.TestCase{
		Description: "TestGetAllProduct_Successfully",
		Request: call.Request{
			URL:    "http://localhost:8080/products",
			Method: http.MethodGet,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `[
				{
					"name":"foo",
					"stock_quantity":0
				},
				{
					"name":"bar",
					"stock_quantity":0
				}
			]`,
			IgnoreBodyFields: []string{
				"0.id", "0.code", "0.created_at", "0.updated_at",
				"1.id", "1.code", "1.created_at", "1.updated_at",
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}