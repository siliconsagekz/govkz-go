package request

import (
	"net/http"
)

func Request(uri string) (*http.Response, error) {
	res, err := http.Get(uri); if err != nil {
		return nil, err
	}
	
	return res, nil
}