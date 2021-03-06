package test

import "net/http"

type (
	Header struct {
		header http.Header
	}
)

func (h *Header) Add(key, val string) {
	h.header.Add(key, val)
}

func (h *Header) Del(key string) {
	h.header.Del(key)
}

func (h *Header) Get(key string) string {
	return h.header.Get(key)
}

func (h *Header) Set(key, val string) {
	h.header.Set(key, val)
}

func (h *Header) All() []string {
	headers := make([]string, len(h.header))
	idx := 0
	for header, _ := range h.header {
		headers[idx] = header
		idx++
	}
	return headers
}

func (h *Header) Object() interface{} {
	return h.header
}

func (h *Header) reset(hdr http.Header) {
	h.header = hdr
}
