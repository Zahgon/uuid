package uuid

import "io"

func New() UUID { _ = "STUB: not implemented"; return *new(UUID) }

func NewString() string { _ = "STUB: not implemented"; return "" }

func NewRandom() (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

func NewRandomFromReader(r io.Reader) (UUID, error) {
	_ = "STUB: not implemented"
	return *new(UUID), nil
}

func newRandomFromPool() (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }
