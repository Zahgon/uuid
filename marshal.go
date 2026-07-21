package uuid

func (uuid UUID) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (uuid *UUID) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

func (uuid UUID) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (uuid *UUID) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }
