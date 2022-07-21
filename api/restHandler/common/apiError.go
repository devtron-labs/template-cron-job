/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package common

import (
	"net/http"
)

//use of writeJsonRespStructured is preferable. it api exists due to historical reason
// err.message is used as internal message for ApiError object in resp
func WriteJsonResp(w http.ResponseWriter, err error, respBody interface{}, status int) {

}

//global response body used across api
type Response struct {
	Code   int         `json:"code,omitempty"`
	Status string      `json:"status,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

func contains(s []*string, e *string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
