package uuid

import (
	"crypto/rand"
	"errors"
	"io"
	"sync"
)

type UUID [16]byte

type Version byte

type Variant byte

const (
	Invalid = Variant(iota)
	RFC4122
	Reserved
	Microsoft
	Future
)

const Standard = RFC4122

const randPoolSize = 16 * 16

var (
	rander      = rand.Reader
	poolEnabled = false
	poolMu      sync.Mutex
	poolPos     = randPoolSize
	pool        [randPoolSize]byte

	ErrInvalidUUIDFormat      = errors.New("invalid UUID format")
	ErrInvalidBracketedFormat = errors.New("invalid bracketed UUID format")
)

type URNPrefixError struct{ prefix string }

func (e URNPrefixError) Error() string { _ = "STUB: not implemented"; return "" }

func (e URNPrefixError) Is(target error) bool { _ = "STUB: not implemented"; return false }

var ErrInvalidURNPrefix = URNPrefixError{}

type invalidLengthError struct{ len int }

func (err invalidLengthError) Error() string { _ = "STUB: not implemented"; return "" }

func (e invalidLengthError) Is(target error) bool { _ = "STUB: not implemented"; return false }

var ErrInvalidLength = invalidLengthError{}

func IsInvalidLengthError(err error) bool { _ = "STUB: not implemented"; return false }

func Parse(s string) (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

func ParseBytes(b []byte) (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

func MustParse(s string) UUID { _ = "STUB: not implemented"; return *new(UUID) }

func FromBytes(b []byte) (uuid UUID, err error) { _ = "STUB: not implemented"; return *new(UUID), nil }

func Must(uuid UUID, err error) UUID { _ = "STUB: not implemented"; return *new(UUID) }

func Validate(s string) error { _ = "STUB: not implemented"; return nil }

func (uuid UUID) String() string { _ = "STUB: not implemented"; return "" }

func (uuid UUID) URN() string { _ = "STUB: not implemented"; return "" }

func encodeHex(dst []byte, uuid UUID) { _ = "STUB: not implemented"; return }

func (uuid UUID) Variant() Variant { _ = "STUB: not implemented"; return *new(Variant) }

func (uuid UUID) Version() Version { _ = "STUB: not implemented"; return *new(Version) }

func (v Version) String() string { _ = "STUB: not implemented"; return "" }

func (v Variant) String() string { _ = "STUB: not implemented"; return "" }

func SetRand(r io.Reader) { _ = "STUB: not implemented"; return }

func EnableRandPool() { _ = "STUB: not implemented"; return }

func DisableRandPool() { _ = "STUB: not implemented"; return }

type UUIDs []UUID

func (uuids UUIDs) Strings() []string { _ = "STUB: not implemented"; return nil }
