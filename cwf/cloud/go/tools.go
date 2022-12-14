//go:build tools
// +build tools

/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

// Package tools is a convention to ensure `go mod tidy` adds development tool
// packages to go.mod / go.sum.

import (
	_ "magma/orc8r/cloud/go/tools/swaggergen"
)
