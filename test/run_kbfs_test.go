// Copyright 2016 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

// +build !dokan,!fuse

package test

import (
	"testing"

	"github.com/keybase/client/go/logger"
)

func createEngine(t testing.TB) Engine {
	log := logger.NewTestLogger(t)
	return &LibKBFS{log: log}
}
