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
	"encoding/base64"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

func generateRandomString(length int) string {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	rand.Read(b)
	randomStr := base64.StdEncoding.EncodeToString(b)
	return randomStr
}

func initSessionStore() {
	secret := generateRandomString(48)
	sessionStore = sessions.NewCookieStore([]byte(secret))
}

func cleanupSession(w http.ResponseWriter, r *http.Request) {

	session, err := sessionStore.Get(r, "gangway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
}
