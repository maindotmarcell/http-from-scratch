package handler

import (
	"fmt"
	"strings"

	"github.com/maindotmarcell/http-from-scratch/internal/constants"
	"github.com/maindotmarcell/http-from-scratch/internal/http"
)

// Create handler functions here

func HandleRoot(req http.Request) string {
	return http.FormatResponse(http.Response{Status: constants.StatusOK})
}

func HandleEcho(req http.Request) string {
	echoStr := strings.TrimPrefix(req.RequestLine.Path, "/echo/")
	res := http.Response{Status: constants.StatusOK,
		Headers: map[string]string{
			"Content-Type":   "text/plain",
			"Content-Length": fmt.Sprintf("%d", len(echoStr)),
		},
		Body: []byte(echoStr)}
	return http.FormatResponse(res)
}
