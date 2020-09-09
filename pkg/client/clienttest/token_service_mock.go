/*
Copyright 2019-2020 vChain, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clienttest

import (
	"github.com/codenotary/immudb/pkg/client"
)

type TokenServiceMock struct {
	client.TokenService
	GetTokenF       func() (string, error)
	SetTokenF       func(database string, token string) error
	IsTokenPresentF func() (bool, error)
	DeleteTokenF    func() error
}

func (ts TokenServiceMock) GetToken() (string, error) {
	return ts.GetTokenF()
}

func (ts TokenServiceMock) SetToken(database string, token string) error {
	return ts.SetTokenF(database, token)
}

func (ts TokenServiceMock) DeleteToken() error {
	return ts.DeleteTokenF()
}

func (ts TokenServiceMock) IsTokenPresent() (bool, error) {
	return ts.IsTokenPresentF()
}

func (ts TokenServiceMock) GetDatabase() (string, error) {
	return "", nil
}

func (ts TokenServiceMock) WithHds(hds client.HomedirService) client.TokenService {
	return ts
}

func (ts TokenServiceMock) WithTokenFileName(tfn string) client.TokenService {
	return ts
}

// DefaultHomedirServiceMock ...
func DefaultTokenServiceMock() *TokenServiceMock {
	return &TokenServiceMock{
		GetTokenF: func() (string, error) {
			return "", nil
		},
		SetTokenF: func(database string, token string) error {
			return nil
		},
		IsTokenPresentF: func() (bool, error) {
			return true, nil
		},
		DeleteTokenF: func() error {
			return nil
		},
	}
}