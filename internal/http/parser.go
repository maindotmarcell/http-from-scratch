package http

import (
	"fmt"
	"strings"
)

// parses the request, returning a Request struct with RequestLine, Headers and Body populated
func ParseRequest(req []byte) Request {
	requestLine, err := parseReqLine(req)
	if err != nil {
		fmt.Println(err)
	}
	headers, err := parseHeaders(req)
	if err != nil {
		fmt.Println(err)
	}
	body := parseBody(req)

	return Request{
		RequestLine: requestLine,
		Headers:     headers,
		Body:        body,
	}
}

func parseReqLine(req []byte) (RequestLine, error) {
	lines := strings.Split(string(req), "\r\n")

	requestLineParts := strings.Split(lines[0], " ")

	return RequestLine{
		Method:  requestLineParts[0],
		Path:    requestLineParts[1],
		Version: requestLineParts[2],
	}, nil
}

func parseHeaders(req []byte) (Headers, error) {
	reqStr := string(req)

	firstLineEnd := strings.Index(reqStr, "\r\n")
	if firstLineEnd == -1 {
		return Headers{}, fmt.Errorf("error parsing or missing headers")
	}

	headersEnd := strings.Index(reqStr, "\r\n\r\n")
	if headersEnd == -1 {
		return Headers{}, fmt.Errorf("error parsing or missing headers")
	}

	// +2 to get after \r\n
	headersSection := reqStr[firstLineEnd+2 : headersEnd]

	headerLines := strings.Split(headersSection, "\r\n")
	headers := Headers{}

	for _, line := range headerLines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.ToLower(strings.TrimSpace(parts[0]))
		value := strings.TrimSpace(parts[1])

		switch key {
		case "host":
			headers.Host = value
		case "user-agent":
			headers.UserAgent = value
		case "accept":
			headers.Accept = value
		}
	}

	return headers, nil
}

func parseBody(req []byte) []byte {
	reqStr := string(req)

	bodyStart := strings.Index(reqStr, "\r\n\r\n")
	if bodyStart == -1 {
		return []byte("")
	}

	return []byte(reqStr[bodyStart+4:])
}
