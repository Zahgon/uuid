package uuid

import (
	"hash"
)

var (
	NameSpaceDNS  = MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	NameSpaceURL  = MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")
	NameSpaceOID  = MustParse("6ba7b812-9dad-11d1-80b4-00c04fd430c8")
	NameSpaceX500 = MustParse("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
	Nil           UUID

	Max = UUID{
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	}
)

func NewHash(h hash.Hash, space UUID, data []byte, version int) UUID {
	_ = "STUB: not implemented"
	return *new(UUID)
}

//nolint:errcheck
//nolint:errcheck

func NewMD5(space UUID, data []byte) UUID { _ = "STUB: not implemented"; return *new(UUID) }

func NewSHA1(space UUID, data []byte) UUID { _ = "STUB: not implemented"; return *new(UUID) }
