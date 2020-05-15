// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package authz

import (
	"errors"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/clevergo/clevergo"
)

// Errors
var (
	ErrUnauthorized = clevergo.NewError(http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	ErrForbidden    = clevergo.NewError(http.StatusForbidden, errors.New("you are not allowed to access this page"))
)

// UserFunc is a function that returns a string which represents the authenticated user.
type UserFunc func(c *clevergo.Context) (id string, err error)

// New returns a middleware with the given enforcer, user function and optional options.
func New(enforcer *casbin.Enforcer, userFunc UserFunc, opts ...Option) clevergo.MiddlewareFunc {
	a := &authorization{
		enforcer: enforcer,
		userFunc: userFunc,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a.Middleware
}

type authorization struct {
	enforcer *casbin.Enforcer
	userFunc UserFunc
	Skipper  clevergo.Skipper
}

func (a *authorization) Middleware(next clevergo.Handle) clevergo.Handle {
	return func(c *clevergo.Context) error {
		if a.Skipper == nil || !a.Skipper(c) {
			id, err := a.userFunc(c)
			if err != nil {
				return err
			}
			if id == "" {
				return ErrUnauthorized
			}
			ok, err := a.enforcer.Enforce(id, c.Request.URL.Path, c.Request.Method)
			if err != nil {
				return err
			}
			if !ok {
				return ErrForbidden
			}
		}
		return next(c)
	}
}
