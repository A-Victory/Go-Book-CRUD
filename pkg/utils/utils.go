package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//The ParseBody func help to parse the body for the create func, to be able to unmarshal the json
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
