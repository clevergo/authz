// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package authz

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"clevergo.tech/clevergo"
	"github.com/casbin/casbin/v2"
	"github.com/stretchr/testify/assert"
)

func testUserFunc(id string, err error) UserFunc {
	return func(c *clevergo.Context) (string, error) {
		return id, err
	}
}

func getTestEnforcer() *casbin.Enforcer {
	e, _ := casbin.NewEnforcer("casbin_model.conf", "casbin_policy.csv")
	return e
}

func TestNew(t *testing.T) {
	var errNoUser = errors.New("user does not exist")
	cases := []struct {
		userID    string
		userErr   error
		method    string
		url       string
		shouldErr bool
		err       error
	}{
		{"foo", nil, http.MethodGet, "/posts", false, nil},
		{"bar", nil, http.MethodGet, "/posts", false, nil},
		{"", errNoUser, http.MethodGet, "/posts", true, errNoUser},
		{"", nil, http.MethodGet, "/posts", true, ErrUnauthorized},
		{"foo", nil, http.MethodPost, "/posts", false, nil},
		{"foo", nil, http.MethodGet, "/posts/1", false, nil},
		{"foo", nil, http.MethodPut, "/posts/2", false, nil},
		{"foo", nil, http.MethodPatch, "/posts/3", false, nil},
		{"foo", nil, http.MethodDelete, "/posts/4", false, nil},
		{"foo", nil, http.MethodGet, "/invalid", true, ErrForbidden},
		{"invalid user", nil, http.MethodGet, "/posts", true, ErrForbidden},
	}
	for _, test := range cases {
		m := New(getTestEnforcer(), testUserFunc(test.userID, test.userErr))
		handle := func(c *clevergo.Context) error {
			return nil
		}
		handle = m(handle)
		w := httptest.NewRecorder()
		ctx := &clevergo.Context{
			Request:  httptest.NewRequest(test.method, test.url, nil),
			Response: w,
		}
		err := handle(ctx)
		if test.shouldErr {
			assert.Equal(t, test.err, err)
			continue
		}
		assert.Nil(t, err)
	}
}

func TestNewWithOption(t *testing.T) {
	m := New(nil, nil, WithSkipper(func(*clevergo.Context) bool {
		return true
	}))
	handled := false
	handle := m(func(c *clevergo.Context) error {
		handled = true
		return nil
	})
	ctx := &clevergo.Context{}
	assert.Nil(t, handle(ctx))
	assert.True(t, handled)
}
