package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParsBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)

	fmt.Println("x")
	if err == nil {
		err := json.Unmarshal([]byte(body), x)

		if err != nil {
			return
		}
	}

}
