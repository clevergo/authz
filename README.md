# CleverGo Authorization Middleware
[![Build Status](https://img.shields.io/travis/clevergo/authz?style=for-the-badge)](https://travis-ci.org/clevergo/authz)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/authz?style=for-the-badge)](https://coveralls.io/github/clevergo/authz?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/authz?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/authz?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/authz)
[![Release](https://img.shields.io/github/release/clevergo/authz.svg?style=for-the-badge)](https://github.com/clevergo/authz/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/month/clevergo.tech/authz&style=for-the-badge)](https://pkg.clevergo.tech/)

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
