// Copyright © 2017 Heptio
// Copyright © 2017 Craig Tracey
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

package main

import (
	"fmt"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	rs := generateRandomString(48)
	rs2 := generateRandomString(48)
	fmt.Println(rs)
	fmt.Println(rs2)
	if rs == "" {
		t.Errorf("Received an empty string")
		return
	}
	if rs == rs2 {
		t.Errorf("Generated the same string two times in a row. String is not random.")
		return
	}
}
