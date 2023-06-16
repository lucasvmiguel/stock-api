package api

import (
	"net/http"
	"testing"

	"github.com/lucasvmiguel/integration"
	"github.com/lucasvmiguel/integration/call"
	"github.com/lucasvmiguel/integration/expect"
)

func TestGet_Successfully(t *testing.T) {
	reload()

	err := integration.Test(&integration.HTTPTestCase{
		Description: "TestGet_Successfully",
		Request: call.Request{
			URL:    "http://localhost:8080/products?limit=2",
			Method: http.MethodGet,
		},
		Response: expect.Response{
			StatusCode: http.StatusOK,
			Body: `{
				"items": [
					{
						"id": "<<PRESENCE>>",
						"name": "foo",
						"code": "<<PRESENCE>>",
						"stock_quantity": 1,
						"created_at": "<<PRESENCE>>",
						"updated_at": "<<PRESENCE>>"
					},
					{
						"id": "<<PRESENCE>>",
						"name": "bar",
						"code": "<<PRESENCE>>",
						"stock_quantity": 2,
						"created_at": "<<PRESENCE>>",
						"updated_at": "<<PRESENCE>>"
					}
				],
				"next_cursor": "<<PRESENCE>>"
			}`,
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
