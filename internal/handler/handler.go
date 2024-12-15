package handler

import (
	"github.com/maindotmarcell/http-from-scratch/internal/constants"
	"github.com/maindotmarcell/http-from-scratch/internal/http"
)

// Create handler functions here

func HandleRoot(req http.Request) string {
	return http.FormatResponse(http.Response{Status: constants.StatusOK})
}
