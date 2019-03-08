// Copyright © 2019 Heptio
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
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

// The CustomCookieStore automatically splits cookies with length greater than maxCookieLength into multiple smaller cookies.
// The motivation is the browsers' 4KB limit on cookies, which for instance causes problems for large id_tokens in azure.

const (
	// Cookies are limited to 4kb including the length of the cookie name,
	// the cookie name can be up to 256 bytes
	maxCookieLength = 3840
)

type CustomCookieStore struct {
	*sessions.CookieStore
}

// Set secureCookie maxLength to an arbitrary (20x4kb) high value since we are no longer limited
func NewCustomCookieStore(keyPairs ...[]byte) *CustomCookieStore {
	cookieStore := sessions.NewCookieStore(keyPairs...)
	for _, codec := range cookieStore.Codecs {
		cookie := codec.(*securecookie.SecureCookie)
		cookie.MaxLength(81920)
	}
	return &CustomCookieStore{cookieStore}
}

func (s *CustomCookieStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	return sessions.GetRegistry(r).Get(s, name)
}

// In contrast to default implementation, the session values can be partitioned into
// multiple cookies.
// The original cookie is split/joined in its encoded form
func (s *CustomCookieStore) New(r *http.Request, name string) (*sessions.Session, error) {
	session := sessions.NewSession(s, name)
	opts := *s.Options
	session.Options = &opts
	session.IsNew = true
	cookie := joinSectionCookies(r, name)
	var err error
	if len(cookie) > 0 {
		err = securecookie.DecodeMulti(name, cookie, &session.Values, s.Codecs...)
		if err == nil {
			session.IsNew = false
		}
	}
	return session, err
}

// If the cookie length is > maxCookieLength, its value is split into multiple cookies
// fitting into the maxCookieLength limit.
// The resulting section cookies get their index appended to the name.
func (s *CustomCookieStore) Save(r *http.Request, w http.ResponseWriter,
	session *sessions.Session) error {

	cookie, err := securecookie.EncodeMulti(session.Name(), session.Values,
		s.Codecs...)
	if err != nil {
		return err
	}

	sectionCookies := splitCookie(cookie)
	// With a singular section the name is unchanged
	if len(sectionCookies) == 1 {
		cookieName := session.Name()
		http.SetCookie(w, sessions.NewCookie(cookieName, sectionCookies[0], session.Options))
		return nil
	}

	for i, value := range sectionCookies {
		cookieName := buildSectionCookieName(session.Name(), i)
		http.SetCookie(w, sessions.NewCookie(cookieName, value, session.Options))
	}
	return nil
}

// joinCookies concatenates the values of all matching cookies and returns the original, encoded cookievalue string.
func joinSectionCookies(r *http.Request, name string) string {

	// Exact match without index means only a single cookie exists
	if c, err := r.Cookie(name); err == nil {
		return c.Value
	}

	var joinedValue string
	for i := 0; true; i++ {
		cookieName := buildSectionCookieName(name, i)
		if c, err := r.Cookie(cookieName); err == nil {
			joinedValue += c.Value
		} else {
			break
		}
	}
	return joinedValue
}

// splitCookie splits the original encoded cookie value into a slice of cookies which
// fit within the 4kb cookie limit indexing the cookies from 0
func splitCookie(cookieValue string) []string {
	var sectionCookies []string
	valueBytes := []byte(cookieValue)

	for len(valueBytes) > 0 {
		length := len(valueBytes)
		if length > maxCookieLength {
			length = maxCookieLength
		}
		sectionCookies = append(sectionCookies, string(valueBytes[:length]))
		valueBytes = valueBytes[length:]
	}
	return sectionCookies
}

func buildSectionCookieName(name string, index int) string {
	return fmt.Sprintf("%s_%d", name, index)
}
