/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

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

package server

import (
	"github.com/codenotary/immudb/pkg/database"
	"github.com/codenotary/immudb/pkg/logger"
)

type Option func(s *srv)

func Host(c string) Option {
	return func(args *srv) {
		args.Host = c
	}
}

func Port(port string) Option {
	return func(args *srv) {
		args.Port = port
	}
}

func Logger(logger logger.Logger) Option {
	return func(args *srv) {
		args.Logger = logger
	}
}

func DatabaseList(dbList database.DatabaseList) Option {
	return func(args *srv) {
		args.dbList = dbList
	}
}
