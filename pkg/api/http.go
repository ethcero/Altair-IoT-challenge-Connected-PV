package api

import (
	"io"
	"net/http"
)

type Request struct {
	Url string
}
type Response struct {
	Body []byte
}

func Get(req Request) (*Response, error) {
	resp, err := http.Get(req.Url)
	if err != nil {
		return &Response{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{}, err
	}
	return &Response{
		body,
	}, nil
}
