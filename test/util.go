//go:build testing

package test

import "os"

func EnvOrDefault(name string, fallback string) string {
	if val, ok := os.LookupEnv(name); !ok {
		return fallback
	} else {
		return val
	}
}
