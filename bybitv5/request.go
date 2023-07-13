package bybitv5

import (
	"fmt"
	"net/http"

	"github.com/msw-x/moon/uhttp"
)

func req[R, T any](c *Client, method string, path string, request any, transform func(R) (T, error), sign bool) (r Response[T]) {
	var perf *uhttp.Performer
	switch method {
	case http.MethodGet:
		perf = c.c.Get(path).Params(request)
	case http.MethodPost:
		perf = c.c.Post(path).Json(request)
	default:
		r.Error = fmt.Errorf("forbidden method: %s", method)
		return
	}
	if sign && c.s != nil {
		if perf.Request.Header == nil {
			perf.Request.Header = make(http.Header)
		}
		switch method {
		case http.MethodGet:
			c.s.HeaderGet(perf.Request.Header, perf.Request.Params)
		case http.MethodPost:
			c.s.HeaderPost(perf.Request.Header, perf.Request.Body)
			perf.Header("Referer", "GinArea").Header("x-referer", "GinArea")
		}
	}
	h := perf.Do()
	if h.Error == nil {
		if h.BodyExists() {
			raw := new(response[R])
			h.Json(raw)
			r.Time = raw.Time
			r.Error = raw.Error()
			if r.Ok() {
				r.Data, r.Error = transform(raw.Result)
			}
		}
		if sign {
			r.SetErrorIfNil(h.HeaderTo(&r.Limit))
		}
	} else {
		r.Error = h.Error
		r.NetError = true
	}
	return
}

func request[R, T any](c *Client, method string, path string, request any, transform func(R) (T, error), sign bool) (r Response[T]) {
	var attempt int
	for {
		r = req(c, method, path, request, transform, sign)
		if r.NetError && c.onNetError != nil {
			if c.onNetError(r.Error, attempt) {
				attempt++
				continue
			}
		}
		break
	}
	return
}

func GetPub[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodGet, path, req, transform, false)
}

func Get[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodGet, path, req, transform, true)
}

func Post[R, T any](c *Client, path string, req any, transform func(R) (T, error)) Response[T] {
	return request(c, http.MethodPost, path, req, transform, true)
}
