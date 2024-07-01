// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

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
