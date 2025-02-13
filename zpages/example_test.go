// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zpages_test

import (
	"log"
	"net/http"
	
	"github.com/gozelle/opencensus/zpages"
)

func Example() {
	// Both /debug/tracez and /debug/rpcz will be served on the default mux.
	zpages.Handle(nil, "/debug")
	log.Fatal(http.ListenAndServe("127.0.0.1:9999", nil))
}
