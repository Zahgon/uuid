package uuid

import (
	"database/sql/driver"
)

var jsonNull = []byte("null")

type NullUUID struct {
	UUID  UUID
	Valid bool
}

func (nu *NullUUID) Scan(value interface{}) error { _ = "STUB: not implemented"; return nil }

func (nu NullUUID) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (nu NullUUID) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (nu *NullUUID) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

func (nu NullUUID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (nu *NullUUID) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

func (nu NullUUID) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (nu *NullUUID) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
