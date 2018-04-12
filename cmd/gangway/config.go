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
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// Config the configuration field for gangway
type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	ClusterName   string   `yaml:"clusterName" envconfig:"cluster_name"`
	AuthorizeURL  string   `yaml:"authorizeURL" envconfig:"authorize_url"`
	TokenURL      string   `yaml:"tokenURL" envconfig:"token_url"`
	ClientID      string   `yaml:"clientID" envconfig:"client_id"`
	ClientSecret  string   `yaml:"clientSecret" envconfig:"client_secret"`
	Audience      string   `yaml:"audience"`
	RedirectURL   string   `yaml:"redirectURL" envconfig:"redirect_url"`
	Scopes        []string `yaml:"scopes"`
	UsernameClaim string   `yaml:"usernameClaim" envconfig:"username_claim"`
	EmailClaim    string   `yaml:"emailClaim" envconfig:"email_claim"`
}

// NewConfig returns a Config struct from serialized config file
func NewConfig(configFile string) (*Config, error) {
	cfg := &Config{
		Host:          "0.0.0.0",
		Port:          8080,
		Scopes:        []string{"openid", "profile", "email", "offline_access"},
		UsernameClaim: "nickname",
		EmailClaim:    "email",
	}

	if configFile != "" {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal([]byte(data), cfg)
		if err != nil {
			return nil, err
		}
	}

	err := envconfig.Process("gangway", cfg)
	if err != nil {
		return nil, err
	}

	err = validateConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	checks := []struct {
		bad    bool
		errMsg string
	}{
		{cfg.AuthorizeURL == "", "no authorizeURL specified"},
		{cfg.TokenURL == "", "no tokenURL specified"},
		{cfg.ClientID == "", "no clientID specified"},
		{cfg.ClientSecret == "", "no clientSecret specified"},
		{cfg.RedirectURL == "", "no redirectURL specified"},
	}

	for _, check := range checks {
		if check.bad {
			return fmt.Errorf("invalid config: %s", check.errMsg)
		}
	}
	return nil
}
