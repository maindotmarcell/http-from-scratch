package http

import (
	"fmt"
	"strings"
)

type Response struct {
	Status  string
	Headers map[string]string
	Body    []byte
}

func FormatResponse(res Response) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("HTTP/1.1 %s\r\n", res.Status))

	for k, v := range res.Headers {
		sb.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	sb.WriteString("\r\n")
	sb.Write(res.Body)

	return sb.String()
}
