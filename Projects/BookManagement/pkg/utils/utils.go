package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// this function will help parse the request body for create Book
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
