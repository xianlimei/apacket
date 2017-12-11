package firstblood

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const (
	HttpResponse       = "HTTP/1.1 200 OK\r\n"
	HttpResponseHeader = "Connection: keep-alive\r\nContent-Encoding: gzip\r\nContent-Type: text/html; charset=utf-8\r\n"
	HttpBody           = "<html><head></head><body></body></html>"
)

var HttpServer = []string{
	"Tengine",
	"nginx/1.10.0",
	"Apache/2.2.21",
}

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

type Http struct {
	methodMap map[string]bool
}

func NewHttp() *Http {
	http := &Http{}
	http.init()
	return http
}

func (http *Http) init() {
	http.methodMap = map[string]bool{
		MethodGet:     true,
		MethodHead:    true,
		MethodPost:    true,
		MethodPut:     true,
		MethodPatch:   true,
		MethodDelete:  true,
		MethodConnect: true,
		MethodOptions: true,
		MethodTrace:   true,
	}
}

func (http *Http) Fingerprint(data []byte) (identify bool, err error) {
	method := make([]byte, 7)
	for idx, d := range data[:8] {
		if d == 0x20 {
			method = data[:idx]
			break
		}
	}
	_, ok := http.methodMap[string(method)]
	if ok {
		identify = true
	}
	return
}

func (http *Http) DisguiserData() (data []byte) {
	server := fmt.Sprintf("Server: %s\r\n", http.getServer())

	ts := time.Now()
	date := fmt.Sprintf("Date: %s\r\n", ts.UTC().Format(time.UnixDate))

	buf := bytes.Buffer{}
	buf.WriteString(HttpResponse)
	buf.WriteString(date)
	buf.WriteString(HttpResponseHeader)
	buf.WriteString(server)
	buf.WriteString(HttpBody)
	data = buf.Bytes()
	return data
}

func (http *Http) getServer() string {
	rand.Seed(time.Now().UnixNano())
	return HttpServer[rand.Intn(len(HttpServer))]
}
