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

package session

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
	"net/http"
)

const salt = "MkmfuPNHnZBBivy0L0aW"

// Session defines a Gangway session
type Session struct {
	Session *CustomCookieStore
}

// New inits a Session with CookieStore
func New(sessionSecurityKey string) *Session {
	return &Session{
		Session: NewCustomCookieStore(generateSessionKeys(sessionSecurityKey)),
	}
}

// generateSessionKeys creates a signed encryption key for the cookie store
func generateSessionKeys(sessionSecurityKey string) ([]byte, []byte) {
	// Take the configured security key and generate 96 bytes of data. This is
	// used as the signing and encryption keys for the cookie store.  For details
	// on the PBKDF2 function: https://en.wikipedia.org/wiki/PBKDF2
	b := pbkdf2.Key(
		[]byte(sessionSecurityKey),
		[]byte(salt),
		4096, 96, sha256.New)

	return b[0:64], b[64:96]
}

// Cleanup removes the current session from the store
func (s *Session) Cleanup(w http.ResponseWriter, r *http.Request, name string) {
	session, err := s.Session.Get(r, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
}
