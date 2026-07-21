//go:build !js
// +build !js

package uuid

import "net"

var interfaces []net.Interface

func getHardwareInterface(name string) (string, []byte) { _ = "STUB: not implemented"; return "", nil }
