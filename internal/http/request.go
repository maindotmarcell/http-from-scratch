package http

type RequestLine struct {
	Method  string
	Path    string
	Version string
}

type Headers struct {
	Host      string
	UserAgent string
	Accept    string
}

type Request struct {
	RequestLine RequestLine
	Headers     Headers
	Body        []byte
}
