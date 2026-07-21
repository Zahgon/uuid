package uuid

import (
	"sync"
)

var (
	nodeMu sync.Mutex
	ifname string
	nodeID [6]byte
	zeroID [6]byte
)

func NodeInterface() string { _ = "STUB: not implemented"; return "" }

func SetNodeInterface(name string) bool { _ = "STUB: not implemented"; return false }

func setNodeInterface(name string) bool { _ = "STUB: not implemented"; return false }

func NodeID() []byte { _ = "STUB: not implemented"; return nil }

func SetNodeID(id []byte) bool { _ = "STUB: not implemented"; return false }

func (uuid UUID) NodeID() []byte { _ = "STUB: not implemented"; return nil }
