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

func TestDownload(t *testing.T) {
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
	t.Log("Should download file with correct paramter given")
	if _, err := Download(urlStr, file, &client); err == nil {
		t.Logf("\t %s Download file success ", succeed)
	} else {
		t.Fatalf("\t %s Download file FAILED", failed)
	}

	// t.Log("Should cause err")

	// if _, err := Download("null", file, &client); err == nil {
	// 	t.Fatalf("\t %s Test failed", failed)
	// } else {
	// 	t.Logf("\t %s, Test succeed", succeed)
	// }
}
