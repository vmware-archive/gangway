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
	"github.com/gorilla/sessions"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJoinSectionCookies(t *testing.T) {
	var originalValue string
	var value string
	cookies := buildRandomCookies(2, 3800, "test_%d")
	buildRequestWithCookies(cookies, func(cookies []*http.Cookie, r *http.Request) {
		for _, c := range cookies {
			originalValue += c.Value
		}
		value = joinSectionCookies(r, "test")
	})
	if value != originalValue {
		t.Errorf("joinSectionCookies value incorrect: \n value: %s \n originalValue: %s", value, originalValue)
	}
}

func TestJoinSectionCookiesSingle(t *testing.T) {
	var originalValue string
	var value string
	cookies := buildRandomCookies(1, 2000, "test_%d")
	buildRequestWithCookies(cookies, func(cookies []*http.Cookie, r *http.Request) {
		for _, c := range cookies {
			originalValue += c.Value
		}
		value = joinSectionCookies(r, "test")
	})
	if value != originalValue {
		t.Errorf("joinSectionCookies value incorrect: \n value: %s \n originalValue: %s", value, originalValue)
	}
}

func TestSplitCookie(t *testing.T) {
	cookieLength := 8000
	originalValue := randStringBytesRmndr(cookieLength)
	sectionCookies := splitCookie(originalValue)
	expectedCount := int(math.Ceil((float64(cookieLength) / maxCookieLength)))
	if len(sectionCookies) != expectedCount {
		t.Errorf("splitCookie count incorrect: \n count: %d \n expectedCount: %d", len(sectionCookies), expectedCount)
	}
	value := strings.Join(sectionCookies, "")
	if value != originalValue {
		t.Errorf("splitCookie value incorrect: \n value: %s \n originalValue: %s", value, originalValue)
	}
}

func TestSplitCookieSingle(t *testing.T) {
	cookieLength := 2000
	originalValue := randStringBytesRmndr(cookieLength)
	sectionCookies := splitCookie(originalValue)
	expectedCount := int(math.Ceil((float64(cookieLength) / maxCookieLength)))
	if len(sectionCookies) != expectedCount {
		t.Errorf("splitCookie count incorrect: \n count: %d \n expectedCount: %d", len(sectionCookies), expectedCount)
	}
}

func TestSplitCookieSize(t *testing.T) {
	cookieLength := 10000
	originalValue := randStringBytesRmndr(cookieLength)
	sectionCookies := splitCookie(originalValue)
	for _, s := range sectionCookies {
		if len(s) > maxCookieLength {
			t.Errorf("sectionCookie length over limit: \n length: %d", len(s))
		}
	}
}

func TestSplitAndJoin(t *testing.T) {
	cookieLength := 10000
	originalValue := randStringBytesRmndr(cookieLength)
	sectionCookies := splitCookie(originalValue)
	cookies := buildCookiesFromValues(sectionCookies, "test_%d")
	var value string
	buildRequestWithCookies(cookies, func(cookies []*http.Cookie, r *http.Request) {
		value = joinSectionCookies(r, "test")
	})
	if value != originalValue {
		t.Errorf("SplitAndJoin value incorrect: \n value: %s \n originalValue: %s", value, originalValue)
	}
}

// Utility

type handleReq func([]*http.Cookie, *http.Request)

func buildRequestWithCookies(cookies []*http.Cookie, fn handleReq) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, cookie := range cookies {
			r.AddCookie(cookie)
		}
		fn(cookies, r)
	}))
	defer ts.Close()
	_, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
}

func buildRandomCookies(cookieCount int, cookieLength int, cookieName string) []*http.Cookie {
	sessionOptions := &sessions.Options{}
	var cookies []*http.Cookie
	for i := 0; i < cookieCount; i++ {
		value := randStringBytesRmndr(cookieLength)
		cookie := sessions.NewCookie(fmt.Sprintf(cookieName, i), value, sessionOptions)
		cookies = append(cookies, cookie)
	}
	return cookies
}

func buildCookiesFromValues(values []string, cookieName string) []*http.Cookie {
	sessionOptions := &sessions.Options{}
	var cookies []*http.Cookie
	for i, value := range values {
		cookie := sessions.NewCookie(fmt.Sprintf(cookieName, i), value, sessionOptions)
		cookies = append(cookies, cookie)
	}
	return cookies
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
