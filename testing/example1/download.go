package example1

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Download(urlStr string, f *os.File, c *http.Client) (int64, error) {
	respond, getErr := c.Get(urlStr)
	if getErr == nil {
		if size, err := io.Copy(f, respond.Body); err == nil {
			return size, nil
		} else {
			return 0, err
		}
	}

	defer respond.Body.Close()
	return 0, getErr
}
func Run() {
	var urlStr string = "http://ipv4.download.thinkbroadband.com/5MB.zip"
	if urlParsed, err := url.Parse(urlStr); err == nil {
		path := urlParsed.Path
		segments := strings.Split(path, "/")
		fileName := segments[len(segments)-1]

		if file, err := os.Create(fileName); err == nil {
			client := http.Client{
				CheckRedirect: func(r *http.Request, via []*http.Request) error {
					r.URL.Opaque = r.URL.Path
					return nil
				},
			}
			if size, err := Download(urlStr, file, &client); err == nil {
				fmt.Printf("Got: %d", size)
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
	}

}
