package example1

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

type TestParams struct {
	name   string
	urlStr string
}

func TestDownload(t *testing.T) {

	testParams := []TestParams{
		{"Good Params", "http://ipv4.download.thinkbroadband.com/5MB.zip"},
		{"Good Params", "http://ipv4.download.thinkbroadband.com/10MB.zip"},
	}

	for _, e := range testParams {
		t.Logf("Running test :  %s", e.name)
		testingFunc := func(t *testing.T) {

			var urlStr string = "http://ipv4.download.thinkbroadband.com/5MB.zip"
			urlParsed, _ := url.Parse(urlStr)
			path := urlParsed.Path
			segments := strings.Split(path, "/")
			fileName := segments[len(segments)-1]

			file, _ := os.Create(fileName)
			client := http.Client{
				CheckRedirect: func(r *http.Request, via []*http.Request) error {
					r.URL.Opaque = r.URL.Path
					return nil
				},
			}
			_, err := Download(e.urlStr, file, &client)
			switch e.name {
			case "Good Params":
				if err == nil {
					t.Logf("Test %s succeed", e.name)
				} else {
					t.Fatalf("Test %s failed", e.name)
				}
			case "Bad Params":
				if err == nil {
					t.Fatalf("Test %s failed", e.name)
				} else {
					t.Logf("Test %s succeed", e.name)
				}

			}

		}
		t.Run(e.name, testingFunc)
	}
}
