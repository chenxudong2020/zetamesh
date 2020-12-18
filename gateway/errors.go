// Copyright 2020 ZetaMesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package gateway

import "github.com/zetamesh/zetamesh/message"

// Error represent a dedicated error type, which contain the API status code
type Error struct {
	Code message.StatusCode
	Err  error
}

// Error implements the error interface
func (e *Error) Error() string {
	return e.Err.Error()
}

// ErrorWithCode returns a error with the specified error message and code
func withcode(err error, code message.StatusCode) error {
	return &Error{
		Code: code,
		Err:  err,
	}
}
