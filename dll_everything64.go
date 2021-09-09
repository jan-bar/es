//go:build windows && amd64
// +build windows,amd64
package es

import _ "embed"

//go:embed Everything64.dll
var everythingDll []byte

var everythingMd5 = []byte("\xa1\xdd\xb6\x98\x1a\xc5\xe0\x80\x55\x4b\xd3\x84\xc8\x69\xf5\xf3")
