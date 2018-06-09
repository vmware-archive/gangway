package main

import "os"

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

func shouldServeTLS(cfg *Config) bool {
	if exists, err := fileExists(cfg.CertFile); !exists || err != nil {
		return false
	} else if exists, err := fileExists(cfg.KeyFile); !exists || err != nil {
		return false
	}

	return true
}
