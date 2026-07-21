package uuid

import (
	"io"
)

func NewV7() (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

func NewV7FromReader(r io.Reader) (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

func makeV7(uuid []byte) { _ = "STUB: not implemented"; return }

var lastV7time int64

const nanoPerMilli = 1000000

func getV7Time() (milli, seq int64) { _ = "STUB: not implemented"; return 0, 0 }
