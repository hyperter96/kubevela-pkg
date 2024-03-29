/*
Copyright 2023 The KubeVela Authors.

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

package providers

const (
	// ParamsKey key for accept params
	ParamsKey = "$params"
	// ReturnsKey key for returned values
	ReturnsKey = "$returns"
)

// Params is the input parameters of a provider.
type Params[T any] struct {
	Params T `json:"$params"`
}

// Returns is the output of a provider.
type Returns[T any] struct {
	Returns T `json:"$returns"`
}
