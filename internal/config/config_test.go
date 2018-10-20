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

package config

import (
	"os"
	"testing"
)

func TestConfigNotFound(t *testing.T) {
	_, err := NewConfig("nonexistentfile")
	if err == nil {
		t.Errorf("Expected config file parsing to file for non-existent config file")
	}
}

func TestEnvionmentOverrides(t *testing.T) {
	os.Setenv("GANGWAY_AUTHORIZE_URL", "https://foo.bar/authorize")
	os.Setenv("GANGWAY_APISERVER_URL", "https://k8s-api.foo.baz")
	os.Setenv("GANGWAY_CLIENT_ID", "foo")
	os.Setenv("GANGWAY_CLIENT_SECRET", "bar")
	os.Setenv("GANGWAY_PORT", "1234")
	os.Setenv("GANGWAY_REDIRECT_URL", "https://foo.baz/callback")
	os.Setenv("GANGWAY_CLUSTER_CA_PATH", "/etc/ssl/certs/ca-certificates.crt")
	os.Setenv("GANGWAY_SESSION_SECURITY_KEY", "testing")
	os.Setenv("GANGWAY_TOKEN_URL", "https://foo.bar/token")
	os.Setenv("GANGWAY_AUDIENCE", "foo")
	os.Setenv("GANGWAY_SCOPES", "groups,sub")
	cfg, err := NewConfig("")
	if err != nil {
		t.Errorf("Failed to test config overrides with error: %s", err)
	}

	if cfg.Port != 1234 {
		t.Errorf("Failed to override config with environment")
	}

	if cfg.Audience != "foo" {
		t.Errorf("Failed to set audience via environment variable. Expected %s but got %s", "foo", cfg.Audience)
	}

	if cfg.Scopes[0] != "groups" || cfg.Scopes[1] != "sub" {
		t.Errorf("Failed to set scopes via environment variable. Expected %s but got %s", "[groups, sub]", cfg.Scopes)
	}
}

func TestGetRootPathPrefix(t *testing.T) {
	tests := map[string]struct {
		path string
		want string
	}{
		"not specified": {
			path: "",
			want: "/",
		},
		"specified": {
			path: "/gangway",
			want: "/gangway",
		},
		"specified default": {
			path: "/",
			want: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cfg := &Config{
				HTTPPath: tc.path,
			}

			got := cfg.GetRootPathPrefix()
			if got != tc.want {
				t.Fatalf("GetRootPathPrefix(): want: %v, got: %v", tc.want, got)
			}
		})
	}
}
