//go:build tools
// +build tools

package main

import (
	_ "github.com/mgechev/revive"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
