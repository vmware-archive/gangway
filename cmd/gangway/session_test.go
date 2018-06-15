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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"

	log "github.com/sirupsen/logrus"
)

func TestGenerateSessionKeys(t *testing.T) {
	cfg.SessionSecurityKey = "testing"

	b1, b2 := generateSessionKeys()

	if len(b1) != 64 || len(b2) != 32 {
		t.Errorf("Wrong byte length's returned")
		return
	}
}

func TestInitSessionStore(t *testing.T) {
	initSessionStore()
	if sessionStore == nil {
		t.Errorf("Session Store is nil. Did not get initialized")
		return
	}

}

func TestCleanupSession(t *testing.T) {

	initSessionStore()
	session := &sessions.Session{}
	// create a test http server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ = sessionStore.Get(r, "gangway")
		cleanupSession(w, r)

	}))
	defer ts.Close()
	_, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	if session.Options.MaxAge != -1 {
		t.Errorf("Session was not reset. Have max age of %d. Should have -1", session.Options.MaxAge)
	}
}
