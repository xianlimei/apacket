package http

import (
	"bytes"
	"encoding/base64"
	//"fmt"
	"github.com/Acey9/apacket/honeypot/core"
	"github.com/Acey9/apacket/logp"
	"math/rand"
	"net"
	"os/exec"
	"strings"
	"time"
	"unicode"
)

const (
	HttpResponse       = "HTTP/1.1 200 OK\r\n"
	HttpResponseHeader = "Connection: keep-alive\r\nContent-Type: text/html; charset=UTF-8\r\nCache-Control: no-cache\r\nPragma: no-cache\r\n"
	HttpBody           = "<html><head><title>Document Error: Unauthorized</title></head><body><h2>Access Error: Unauthorized</h2><p>Access to this document requires a User ID</p>CGI process file does not exist</body></html>"
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

	MethodRTSPDescribe     = "DESCRIBE"
	MethodRTSPSetup        = "SETUP"
	MethodRTSPPlay         = "PLAY"
	MethodRTSPTeardown     = "TEARDOWN"
	MethodRTSPPause        = "PAUSE"
	MethodRTSPRecord       = "RECORD"
	MethodRTSPAnnounce     = "ANNOUNCE"
	MethodRTSPSetParameter = "SET_PARAMETER"
	MethodRTSPGetParameter = "GET_PARAMETER"
	MethodRTSPRedirect     = "REDIRECT"
)

const (
	CmdHttpResponse = "fb_http_response"
	PtypeHTTP       = "http"
	PtypeHTTPS      = "https"
)

type parserState uint8

var (
	transferEncodingChunked = []byte("chunked")

	constCRLF            = []byte("\r\n")
	nameContentLength    = []byte("content-length")
	nameContentType      = []byte("content-type")
	nameTransferEncoding = []byte("transfer-encoding")
	nameConnection       = []byte("connection")
)

const (
	stateStart parserState = iota
	stateFLine
	stateHeaders
	stateBody
	stateBodyChunkedStart
	stateBodyChunked
	stateBodyChunkedWaitFinalCRLF
)

var methodMap = map[string]bool{
	MethodGet:              true,
	MethodHead:             true,
	MethodPost:             true,
	MethodPut:              true,
	MethodPatch:            true,
	MethodDelete:           true,
	MethodConnect:          true,
	MethodTrace:            true,
	MethodRTSPDescribe:     true,
	MethodRTSPSetup:        true,
	MethodRTSPPlay:         true,
	MethodRTSPTeardown:     true,
	MethodRTSPPause:        true,
	MethodRTSPRecord:       true,
	MethodRTSPAnnounce:     true,
	MethodRTSPSetParameter: true,
	MethodRTSPGetParameter: true,
	MethodRTSPRedirect:     true,
}

var protocolMap = map[string]bool{
	"HTTP": true,
	"RTSP": true,
	"MCTP": true,
	"SIP":  true,
}

type HTTP struct {
}

type HTTPMsg struct {
	parseOffset      int
	parseState       parserState
	Method           string            `json:"method,omitempty"`
	RequestURI       string            `json:"uri,omitempty"`
	Version          string            `json:"version,omitempty"`
	Headers          map[string]string `json:"-"`
	Body             string            `json:"-"`
	Payload          []byte            `json:"-"`
	contentLength    int
	hasContentLength bool
	contentType      string
	transferEncoding string
	connection       string
}

func NewHTTP() *HTTP {
	http := &HTTP{}
	return http
}

func (http *HTTP) ParseHttpLine(request []byte) (method, uri, version string) {
	i := bytes.Index(request, []byte("\r\n"))
	if i == -1 {
		afterMethodIdx := bytes.IndexFunc(request, unicode.IsSpace)
		afterRequestURIIdx := bytes.LastIndexFunc(request, unicode.IsSpace)
		if afterMethodIdx == -1 || afterRequestURIIdx == -1 || afterMethodIdx == afterRequestURIIdx {
			return "", "", ""
		}
		method = string(request[:afterMethodIdx])
		uri = string(request[afterMethodIdx+1:])
		return "", "", ""
	}
	fline := request[:i]
	afterMethodIdx := bytes.IndexFunc(fline, unicode.IsSpace)
	afterRequestURIIdx := bytes.LastIndexFunc(fline, unicode.IsSpace)
	if afterMethodIdx == -1 || afterRequestURIIdx == -1 || afterMethodIdx == afterRequestURIIdx {
		return
	}
	method = string(fline[:afterMethodIdx])
	uri = string(fline[afterMethodIdx+1 : afterRequestURIIdx])

	version = string(fline[afterRequestURIIdx+1:])
	protoVersion := strings.Split(version, "/")
	if len(protoVersion) != 2 {
		return "", "", ""
	}
	proto := strings.ToUpper(protoVersion[0])
	_, ok := protocolMap[proto]
	if !ok {
		return "", "", ""
	}
	return
}

func (http *HTTP) Fingerprint(request []byte, tlsTag bool) (identify bool, ptype string, err error) {
	ptype = PtypeHTTP
	if tlsTag {
		ptype = PtypeHTTPS
	}
	method, uri, version := http.ParseHttpLine(request)
	/*
		_, ok := methodMap[method]
		if ok {
			identify = true
			return
		}*/
	if method != "" && uri != "" && version != "" {
		logp.Debug("HTTP", "http.Fingerprint: method:%s, uri:%s, version:%s", method, uri, version)
		identify = true
	}
	return
}

func (http *HTTP) Parser(remoteAddr, localAddr string, request []byte, ptype string, tls bool) (response *core.Applayer) {
	var reqAddr string
	httpMSG := &HTTPMsg{Payload: request}
	httpMSG.parse()

	response, err := core.NewApplayer(remoteAddr, localAddr, ptype, core.TransportTCP, request, tls, httpMSG)
	if err != nil {
		return
	}

	hostPort, ok := httpMSG.Headers["host"]
	if !ok {
		return
	}

	host, port, ipv, err := core.GetIPPort(hostPort)
	if err != nil {
		return
	}

	if port == 0 {
		return
	}

	if ipv == core.IPv4 {
		reqAddr = response.IP4.Dip
	} else if ipv == core.IPv6 {
		reqAddr = response.IP6.Dip
	} else {
		return
	}

	if reqAddr == host || core.IntranetIP(host) {
		response.TCP.Dport = port
	} else {
		//May need to modify 24 to 32
		_, reqAddrNet, err := net.ParseCIDR(reqAddr + "/24")
		if err != nil {
			return
		}
		_, hostNet, err := net.ParseCIDR(host + "/24")
		if err != nil {
			return
		}
		if reqAddrNet.String() == hostNet.String() {
			response.TCP.Dport = port
		}

	}
	return
}

func (http *HTTP) DisguiserResponse(request []byte) (response []byte) {
	var out bytes.Buffer
	str := base64.StdEncoding.EncodeToString(request)
	cmd := exec.Command(CmdHttpResponse, str)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		logp.Err("http.DisguiserResponse.cmd:%v", err)
		return
	}
	response = out.Bytes()
	/*
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

		response = buf.Bytes()
	*/
	return
}

func (http *HTTP) getServer() string {
	rand.Seed(time.Now().UnixNano())
	return HttpServer[rand.Intn(len(HttpServer))]
}

func (http *HTTP) getAuth() string {
	rand.Seed(time.Now().UnixNano())
	return Authenticate[rand.Intn(len(Authenticate))]
}

func (http *HTTPMsg) parseHTTPLine() (cont, ok, complete bool) {
	i := bytes.Index(http.Payload[http.parseOffset:], []byte("\r\n"))
	if i == -1 {
		return false, false, false
	}
	fline := http.Payload[http.parseOffset:i]
	if len(fline) < 8 {
		return false, false, false
	}
	afterMethodIdx := bytes.IndexFunc(fline, unicode.IsSpace)
	afterRequestURIIdx := bytes.LastIndexFunc(fline, unicode.IsSpace)

	// Make sure we have the VERB + URI + HTTP_VERSION
	if afterMethodIdx == -1 || afterRequestURIIdx == -1 || afterMethodIdx == afterRequestURIIdx {
		return false, false, false
	}

	http.Method = string(fline[:afterMethodIdx])
	http.RequestURI = string(fline[afterMethodIdx+1 : afterRequestURIIdx])
	http.Version = string(fline[afterRequestURIIdx+1:])

	http.parseOffset = i + 2
	http.parseState = stateHeaders
	return true, true, true
}

func (http *HTTPMsg) parseHeader(data []byte) (bool, bool, int) {
	if http.Headers == nil {
		http.Headers = make(map[string]string)
	}

	i := bytes.Index(data, []byte(":"))
	if i == -1 {
		// Expected \":\" in headers. Assuming incomplete"
		return true, false, 0
	}
	for p := i + 1; p < len(data); {
		q := bytes.Index(data[p:], constCRLF)
		if q == -1 {
			return true, false, 0
		}
		p += q
		if len(data) > p && (data[p+1] == ' ' || data[p+1] == '\t') {
			p = p + 2
		} else {
			var headerNameBuf [140]byte
			headerName := core.ToLower(headerNameBuf[:], data[:i])
			headerVal := core.Trim(data[i+1 : p])

			if bytes.Equal(headerName, nameContentLength) {
				http.contentLength, _ = core.ParseInt(headerVal)
				http.hasContentLength = true
			} else if bytes.Equal(headerName, nameContentType) {
				http.contentType = string(headerVal)
			} else if bytes.Equal(headerName, nameTransferEncoding) {
				http.transferEncoding = string(headerVal)
			} else if bytes.Equal(headerName, nameConnection) {
				http.connection = string(headerVal)
			}

			if val, ok := http.Headers[string(headerName)]; ok {
				composed := make([]byte, len(val)+len(headerVal)+2)
				off := copy(composed, val)
				off = copy(composed[off:], []byte(", "))
				copy(composed[off:], headerVal)

				http.Headers[string(headerName)] = string(composed)
			} else {
				http.Headers[string(headerName)] = string(headerVal)
			}

			return true, true, p + 2
		}

	}
	return true, false, len(data)
}

func (http *HTTPMsg) parseHeaders() (cont, ok, complete bool) {
	if len(http.Payload)-http.parseOffset >= 2 && bytes.Equal(http.Payload[http.parseOffset:http.parseOffset+2], []byte("\r\n")) {
		http.parseOffset += 2

		if bytes.Equal([]byte(http.transferEncoding), transferEncodingChunked) {
			// support for HTTP/1.1 Chunked transfer
			// Transfer-Encoding overrides the Content-Length
			//s.parseState = stateBodyChunkedStart
			//return true, true, true
			//TODO
		}
		if http.contentLength == 0 {
			// Ignore body for request that contains a message body but not a Content-Length
			return false, true, true
		}
		http.parseState = stateBody

	} else {
		ok, hfcomplete, offset := http.parseHeader(http.Payload[http.parseOffset:])
		if !ok {
			return false, false, false
		}
		if !hfcomplete {
			return false, true, false
		}
		http.parseOffset += offset
	}
	return true, true, true
}

func (http *HTTPMsg) parseBody() (ok, complete bool) {
	if http.parseOffset == len(http.Payload) || http.parseOffset+http.contentLength > len(http.Payload) {
		return true, true
	}
	http.Body = string(http.Payload[http.parseOffset : http.parseOffset+http.contentLength])
	return true, true
}

func (http *HTTPMsg) parse() (bool, bool) {
	for http.parseOffset < len(http.Payload) {
		switch http.parseState {
		case stateStart:
			if cont, ok, complete := http.parseHTTPLine(); !cont {
				return ok, complete
			}
		case stateHeaders:
			if cont, ok, complete := http.parseHeaders(); !cont {
				return ok, complete
			}
		case stateBody:
			return http.parseBody()
		}
	}
	return true, true
}
