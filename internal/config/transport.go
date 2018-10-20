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

package config

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

// TransportConfig describes a configured httpClient
type TransportConfig struct {
	HTTPClient *http.Client
}

// NewTransportConfig returns a TransportConfig with configured httpClient
func NewTransportConfig(trustedCAPath string) *TransportConfig {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	if trustedCAPath != "" {
		// Read in the cert file
		certs, err := ioutil.ReadFile(trustedCAPath)
		if err != nil {
			log.Fatalf("Failed to append %q to RootCAs: %v", trustedCAPath, err)
		}

		// Append our cert to the system pool
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			log.Println("No certs appended, using system certs only")
		}
	}

	// Trust the augmented cert pool in our client
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAs,
			},
		},
	}

	return &TransportConfig{
		HTTPClient: httpClient,
	}
}
