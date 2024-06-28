package http_plugin

import (
	"context"
	"io"
	"net/http"
)

type Plugin interface {
	Init() error
	Close() error
	Handle(ctx context.Context, req interface{}, resp interface{}) error
}

type Request interface {
	GetHeader(key string) string
	GetBody() ([]byte, error)
}

type Response interface {
	SetHeader(key, value string)
	SetBody([]byte) error
}

type httpResponse struct {
	resp http.ResponseWriter
}

func NewHTTPResponse(resp http.ResponseWriter) Response {
	return &httpResponse{resp: resp}
}

func (h *httpResponse) SetHeader(key, value string) {
	h.resp.Header().Set(key, value)
}

func (h *httpResponse) SetBody(body []byte) error {
	_, err := h.resp.Write(body)
	if err != nil {
		return err
	}
	return nil
}

type httpRequest struct {
	req *http.Request
}

func NewHTTPRequest(req *http.Request) Request {
	return &httpRequest{req: req}
}

func (r *httpRequest) GetHeader(key string) string {
	return r.req.Header.Get(key)
}

func (r *httpRequest) GetBody() ([]byte, error) {
	return io.ReadAll(r.req.Body)
}
