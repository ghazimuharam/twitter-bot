package resolver

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func GetRedirectURL(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	cl := http.Client{
		Timeout: 2 * time.Second,
	}
	var lastUrlQuery string
	cl.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > 10 {
			return errors.New("too many redirects")
		}
		lastUrlQuery = req.URL.String()
		return nil
	}

	_, err = cl.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to resolve url: %v", err)
	}

	return lastUrlQuery, nil
}
