// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package authz

import (
	"testing"

	"github.com/clevergo/clevergo"
	"github.com/stretchr/testify/assert"
)

func TestWithSkipper(t *testing.T) {
	a := &authorization{}
	skipped := false
	skipper := func(c *clevergo.Context) bool {
		skipped = true
		return true
	}
	WithSkipper(skipper)(a)
	a.Skipper(nil)
	assert.True(t, skipped)
}
