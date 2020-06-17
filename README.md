# CleverGo Authorization Middleware
[![Build Status](https://travis-ci.org/clevergo/authz.svg?branch=master)](https://travis-ci.org/clevergo/authz)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/authz/badge.svg?branch=master)](https://coveralls.io/github/clevergo/authz?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/clevergo.tech/authz?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/authz)](https://goreportcard.com/report/github.com/clevergo/authz)
[![Release](https://img.shields.io/github/release/clevergo/authz.svg?style=flat-square)](https://github.com/clevergo/authz/releases)

## Usage

```go
import (
    "clevergo.tech/authz"
    "clevergo.tech/clevergo"
)
```

```go
enforcer, _ := casbin.NewEnforcer("casbin_model.conf", "casbin_policy.csv")
userFunc := func(c *clevergo.Context) (id string, err error) {
    // returns the authenticated user ID.
    return
}
app := clevergo.New()
app.Use(authz.New(enforcer, userFunc))
```
