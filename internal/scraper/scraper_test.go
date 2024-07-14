package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockHTML = `
  <html>
    <head>
      <title>UK Charts top 100</title>
    </head>
    <body>
      <div class="chart-content">
        <div class="chart-items">
          <div class="chart-item">
            <div class="description">
              <div class="chart-name">
                <span>Song1</span>
              </div>
              <div class="chart-artist">
                <span>Artist1</span>
              <div>
            </div>
          </div>

          <div class="chart-item">
            <div class="description">
              <div class="chart-name">
                <span>Song2</span>
              </div>
              <div class="chart-artist">
                <span>Artist2</span>
              <div>
            </div>
          </div>

          <div class="chart-item">
            <div class="description">
              <div class="chart-name">
                <span>Song3</span>
              </div>
              <div class="chart-artist">
                <span>Artist3</span>
              <div>
            </div>
          </div>
        </div>
      </div>
    </body>
  </html>
`

func TestScrapeSongs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(mockHTML))
		},
	))

	defer server.Close()

	songs, err := ScrapeSongs(server.URL)
	assert.NoError(t, err)
	assert.Len(t, songs, 3)
}
