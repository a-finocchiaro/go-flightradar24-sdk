/*
Forwards requests to endpoints and returns the result.
*/
package webrequest

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func SendRequest(endpoint string) (bytes []byte, err error) {
	res, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	// verify status code is correct
	if res.StatusCode != 200 {
		msg := fmt.Sprintf("HTTP status %s returned from endpoint %s", res.Status, endpoint)
		return nil, errors.New(msg)
	}

	// unpack the bytes and check for an error
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
