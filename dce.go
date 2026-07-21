package uuid

type Domain byte

const (
	Person = Domain(0)
	Group  = Domain(1)
	Org    = Domain(2)
)

func NewDCESecurity(domain Domain, id uint32) (UUID, error) {
	_ = "STUB: not implemented"
	return *new(UUID), nil
}

func NewDCEPerson() (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

func NewDCEGroup() (UUID, error) { _ = "STUB: not implemented"; return *new(UUID), nil }

func (uuid UUID) Domain() Domain { _ = "STUB: not implemented"; return *new(Domain) }

func (uuid UUID) ID() uint32 { _ = "STUB: not implemented"; return 0 }

func (d Domain) String() string { _ = "STUB: not implemented"; return "" }
