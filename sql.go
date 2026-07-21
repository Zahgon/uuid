package uuid

import (
	"database/sql/driver"
)

func (uuid *UUID) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

func (uuid UUID) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}
