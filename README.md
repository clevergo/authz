# CleverGo Authorization Middleware
[![Build Status](https://travis-ci.org/clevergo/authz.svg?branch=master)](https://travis-ci.org/clevergo/authz)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/authz/badge.svg?branch=master)](https://coveralls.io/github/clevergo/authz?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/authz)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/authz)](https://goreportcard.com/report/github.com/clevergo/authz)
[![Release](https://img.shields.io/github/release/clevergo/authz.svg?style=flat-square)](https://github.com/clevergo/authz/releases)
[![Sourcegraph](https://sourcegraph.com/github.com/clevergo/authz/-/badge.svg)](https://sourcegraph.com/github.com/clevergo/authz?badge)


## Usage

```go
import "github.com/clevergo/authz"
```

```go
enforcer, _ := casbin.NewEnforcer("casbin_model.conf", "casbin_policy.csv")
userFunc := func(c *clevergo.Context) (id string, err error) {
    // returns the authenticated user ID.
    return
}
router.Use(authz.New(enforcer, userFunc))
```
