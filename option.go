// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package authz

import "clevergo.tech/clevergo"

// Option is a function that apply on authorization middleware.
type Option func(*authorization)

// WithSkipper is an option that set authorization's skipper.
func WithSkipper(skipper clevergo.Skipper) Option {
	return func(a *authorization) {
		a.skipper = skipper
	}
}
