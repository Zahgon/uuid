package uuid

import (
	"sync"
	"time"
)

type Time int64

const (
	lillian    = 2299160
	unix       = 2440587
	epoch      = unix - lillian
	g1582      = epoch * 86400
	g1582ns100 = g1582 * 10000000
)

var (
	timeMu   sync.Mutex
	lasttime uint64
	clockSeq uint16

	timeNow = time.Now
)

func (t Time) UnixTime() (sec, nsec int64) { _ = "STUB: not implemented"; return 0, 0 }

func GetTime() (Time, uint16, error) { _ = "STUB: not implemented"; return *new(Time), 0, nil }

func getTime(customTime *time.Time) (Time, uint16, error) {
	_ = "STUB: not implemented"
	return *new(Time), 0, nil
}

func ClockSequence() int { _ = "STUB: not implemented"; return 0 }

func clockSequence() int { _ = "STUB: not implemented"; return 0 }

func SetClockSequence(seq int) { _ = "STUB: not implemented"; return }

func setClockSequence(seq int) { _ = "STUB: not implemented"; return }

func (uuid UUID) Time() Time { _ = "STUB: not implemented"; return *new(Time) }

func (uuid UUID) ClockSequence() int { _ = "STUB: not implemented"; return 0 }
