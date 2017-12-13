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

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	ClusterName   string   `yaml:"cluster_name" envconfig:"cluster_name"`
	AuthorizeURL  string   `yaml:"authorize_url" envconfig:"authorize_url"`
	TokenURL      string   `yaml:"token_url" envconfig:"token_url"`
	ClientID      string   `yaml:"client_id" envconfig:"client_id"`
	ClientSecret  string   `yaml:"client_secret" envconfig:"client_secret"`
	Audience      string   `yaml:"audience"`
	RedirectURL   string   `yaml:"redirect_url" envconfig:"redirect_url"`
	Scopes        []string `yaml:"scopes"`
	UsernameClaim string   `yaml:"username_claim" envconfig:"username_claim"`
	EmailClaim    string   `yaml:"email_claim" envconfig:"email_claim"`
}

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
		{cfg.AuthorizeURL == "", "no authorize_url specified"},
		{cfg.TokenURL == "", "no token_url specified"},
		{cfg.ClientID == "", "no client_id specified"},
		{cfg.ClientSecret == "", "no client_secret specified"},
		{cfg.RedirectURL == "", "no redirect_url specified"},
	}

	for _, check := range checks {
		if check.bad {
			return fmt.Errorf("invalid config: %s", check.errMsg)
		}
	}
	return nil
}
