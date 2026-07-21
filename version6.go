package uuid

import (
	"time"
)

func NewV6() (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

func NewV6WithTime(customTime *time.Time) (UUID, error) {
	_ = "STUB: not implemented"
	return *new(UUID), nil
}

func generateV6(now Time, seq uint16) UUID { _ = "STUB: not implemented"; return *new(UUID) }
