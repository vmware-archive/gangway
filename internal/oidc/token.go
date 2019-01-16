// Copyright © 2018 Heptio
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

package oidc

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

// OAuth2Token is an interface which is used when exchanging an id_token for an access token
type OAuth2Token interface {
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
}

// Token is an implementation of OAuth2Token Interface
type Token struct {
	OAuth2Cfg *oauth2.Config
}

// ParseToken returns a jwt token from an idToken, returns error if it cannot parse
func ParseToken(idToken, clientSecret string) (*jwt.Token, error) {
	token, _ := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(clientSecret), nil
	})

	return token, nil
}

// Exchange takes an oauth2 auth token and exchanges for an id_token
func (t *Token) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return t.OAuth2Cfg.Exchange(ctx, code)
}
