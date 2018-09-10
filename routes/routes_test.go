package routes

import (
	"bytes"
	"github.com/franela/goblin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func Test(t *testing.T) {
	r := NewRouter()
	g := goblin.Goblin(t)

	// Test the index endpoint
	g.Describe("Index Endpoint", func() {
		g.It("Should return 200", func() {
			w := performRequest(r, "GET", "/", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
		})
	})
	// Test the GetProduce endpoint
	g.Describe("Get Produce Endpoint", func() {
		g.It("Should return 200", func() {
			w := performRequest(r, "GET", "/GetProduce", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
			g.Assert(w.Body.String()).Equal("[{\"Name\":\"Lettuce\",\"ProduceCode\":\"A12T-4GH7-QPL9-3N4M\",\"UnitPrice\":3.46},{\"Name\":\"Peach\",\"ProduceCode\":\"E5T6-9UI3-TH15-QR88\",\"UnitPrice\":2.99},{\"Name\":\"Green Pepper\",\"ProduceCode\":\"YRT6-72AS-K736-L4AR\",\"UnitPrice\":0.79},{\"Name\":\"Gala Apple\",\"ProduceCode\":\"TQ4C-VV6T-75ZX-1RMR\",\"UnitPrice\":3.59}]")
		})
	})
	// Test the GetProduceByID endpoint
	g.Describe("Get Produce By ID Endpoint", func() {
		g.It("Should return 200", func() {
			w := performRequest(r, "GET", "/GetProduceByID?ProduceCode=A12T-4GH7-QPL9-3N4M", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
		})
	})
	// Test the PostProduce endpoint
	g.Describe("PostProduce Endpoint", func() {
		g.It("Should should return 200", func() {
			body := []byte(`{"name":"Pineapple","produceCode":"A23K-4GH7-QPL9-1B2U","unitPrice":4.26}`)
			w := performRequest(r, "POST", "/PostProduce", bytes.NewReader(body))
			g.Assert(w.Code).Equal(http.StatusOK)
		})
	})
	// Test the GetProduce endpoint after adding pineapple.
	g.Describe("Get Produce Endpoint", func() {
		g.It("Should return 200", func() {
			w := performRequest(r, "GET", "/GetProduce", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
			g.Assert(w.Body.String()).Equal("[{\"Name\":\"Lettuce\",\"ProduceCode\":\"A12T-4GH7-QPL9-3N4M\",\"UnitPrice\":3.46},{\"Name\":\"Peach\",\"ProduceCode\":\"E5T6-9UI3-TH15-QR88\",\"UnitPrice\":2.99},{\"Name\":\"Green Pepper\",\"ProduceCode\":\"YRT6-72AS-K736-L4AR\",\"UnitPrice\":0.79},{\"Name\":\"Gala Apple\",\"ProduceCode\":\"TQ4C-VV6T-75ZX-1RMR\",\"UnitPrice\":3.59},{\"Name\":\"Pineapple\",\"ProduceCode\":\"A23K-4GH7-QPL9-1B2U\",\"UnitPrice\":4.26}]")
		})
	})
	//Test the DeleteProduce endpoint
	g.Describe("DeleteProduce Endpoint", func() {
		g.It("Should should return 200", func() {
			w := performRequest(r, "DELETE", "/DeleteProduce?ProduceCode=A23K-4GH7-QPL9-1B2U", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
		})
	})
	// Test the GetProduce endpoint after deleting pineapple.
	g.Describe("Get Produce Endpoint", func() {
		g.It("Should return 200", func() {
			w := performRequest(r, "GET", "/GetProduce", nil)
			g.Assert(w.Code).Equal(http.StatusOK)
			g.Assert(w.Body.String()).Equal("[{\"Name\":\"Lettuce\",\"ProduceCode\":\"A12T-4GH7-QPL9-3N4M\",\"UnitPrice\":3.46},{\"Name\":\"Peach\",\"ProduceCode\":\"E5T6-9UI3-TH15-QR88\",\"UnitPrice\":2.99},{\"Name\":\"Green Pepper\",\"ProduceCode\":\"YRT6-72AS-K736-L4AR\",\"UnitPrice\":0.79},{\"Name\":\"Gala Apple\",\"ProduceCode\":\"TQ4C-VV6T-75ZX-1RMR\",\"UnitPrice\":3.59}]")
		})
	})
}
