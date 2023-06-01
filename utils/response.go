package utils

import (
	"net/http"
)

func EnsureReaderClosed(response *http.Response) {
	if response != nil {
		response.Body.Close()
	}
}
