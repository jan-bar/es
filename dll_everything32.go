//go:build windows && 386
// +build windows,386
package es

import _ "embed"

//go:embed Everything32.dll
var everythingDll []byte

var everythingMd5 = []byte("\x97\xed\xa9\xe4\x69\xc1\x9f\x1e\x32\x8a\x27\xd9\x94\x56\xe9\x73")
