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
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

var cfg *Config
var oauth2Cfg *oauth2.Config
var sessionStore *sessions.CookieStore

func main() {

	cfgFile := flag.String("config", "", "The config file to use.")
	flag.Parse()

	var err error
	cfg, err = NewConfig(*cfgFile)
	if err != nil {
		log.Errorf("Could not parse config file: %s", err)
		os.Exit(1)
	}

	log.Infof("Running with config %+v", cfg)

	oauth2Cfg = &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       cfg.Scopes,
		Endpoint: oauth2.Endpoint{
			cfg.AuthorizeURL,
			cfg.TokenURL,
		},
	}

	initSessionStore()

	loginRequiredHandlers := alice.New(loginRequired)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/callback", callbackHandler)

	// middleware'd routes
	http.Handle("/logout", loginRequiredHandlers.ThenFunc(logoutHandler))
	http.Handle("/commandline", loginRequiredHandlers.ThenFunc(commandlineHandler))

	bindAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	err = http.ListenAndServe(bindAddr, nil)
	if err != nil {
		log.Errorf("Service failed to start: %s", err)
		os.Exit(1)
	}
}
