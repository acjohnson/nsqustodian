/*
Copyright © 2024 Aaron Johnson <acjohnson@pcdomain.com>

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package nsqadmin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func NsqAdminCall(nsqadminAddr string, httpHeaders string, payload []byte, url string, method string) (string, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return "Failed to create http request: ", err
	}
	headerStrings := strings.Split(httpHeaders, ",")

	for _, headerString := range headerStrings {
		// Split the header string into key and value
		keyValue := strings.SplitN(headerString, ":", 2)
		if len(keyValue) == 2 {
			// Set the header in the request header
			req.Header.Set(keyValue[0], keyValue[1])
		}
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request and check the response status code
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Error in http response: ", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response body: ", err
	}

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		return fmt.Sprintf("Failed to make NSQ admin call, status %s, response: %s", string(resp.StatusCode), string(bodyBytes)), err
	}

	return string(body), nil
}
