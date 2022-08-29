package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// The ParseBody func help to parse the body for the create func, to be able to unmarshal the json
func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}
