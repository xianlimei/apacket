package firstblood

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

const (
	HttpResponse       = "HTTP/1.1 200 OK\r\n"
	HttpResponseHeader = "Connection: keep-alive\r\nContent-Type: text/html; charset=UTF-8\r\nCache-Control: no-cache\r\nPragma: no-cache\r\n"
	HttpBody           = "<html><head><title>Document Error: Unauthorized</title></head><body><h2>Access Error: Unauthorized</h2><p>Access to this document requires a User ID</p></body></html>"
)

var HttpServer = []string{
	"Tengine",
	"nginx/1.10.0",
	"Apache/2.2.21",
	"gSOAP/2.7",
	"GoAhead-Webs",
	"GoAhead-http",
	"RomPager/4.07 UPnP/1.0",
	"lighttpd/1.4.34",
	"Lighttpd/1.4.28",
	"lighttpd/1.4.31",
	"Linux/2.x UPnP/1.0 Avtech/1.0",
	"P-660HW-T1 v3",
	"U S Software Web Server",
	"Netwave IP Camera",
}

var Authenticate = []string{
	`WWW-Authenticate: Basic realm="iPEX Internet Cafe"`,
	`WWW-Authenticate: Digest realm="IgdAuthentication", domain="/", nonce="N2UyNjgxMjA6NjQ1MWZiOTA6IDJlNjI5NDA=", qop="auth", algorithm=MD5`,
	`WWW-Authenticate: Basic realm="NETGEAR DGN1000 "`,
	`WWW-Authenticate: Digest realm="GoAhead", domain=":81",qop="auth", nonce="405448722b302b85aa6ef2b444ea6b5c", opaque="5ccc069c403ebaf9f0171e9517f40e41",algorithm="MD5", stale="FALSE"`,
	`WWW-Authenticate: Basic realm="HomeHub"`,
	`WWW-Authenticate: Basic realm="MOBOTIX Camera User"`,
	`Authorization: Basic aHR0cHdhdGNoOmY=`,
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

func (http *Http) Fingerprint(request []byte) (identify bool, err error) {
	method := make([]byte, 7)
	for idx, d := range request[:8] {
		if d == 0x20 {
			method = request[:idx]
			break
		}
	}
	_, ok := http.methodMap[string(method)]
	if ok {
		identify = true
	}
	return
}

func (http *Http) DisguiserResponse(request []byte) (reponse []byte) {
	server := fmt.Sprintf("Server: %s\r\n", http.getServer())

	ts := time.Now()
	date := fmt.Sprintf("Date: %s\r\n", ts.UTC().Format(time.UnixDate))

	auth := fmt.Sprintf("%s\r\n", http.getAuth())

	buf := bytes.Buffer{}
	buf.WriteString(HttpResponse)
	buf.WriteString(date)
	buf.WriteString(HttpResponseHeader)
	buf.WriteString(server)
	buf.WriteString(auth)

	buf.WriteString("\r\n")

	buf.WriteString(HttpBody)

	reponse = buf.Bytes()
	return
}

func (http *Http) getServer() string {
	rand.Seed(time.Now().UnixNano())
	return HttpServer[rand.Intn(len(HttpServer))]
}

func (http *Http) getAuth() string {
	rand.Seed(time.Now().UnixNano())
	return Authenticate[rand.Intn(len(Authenticate))]
}
