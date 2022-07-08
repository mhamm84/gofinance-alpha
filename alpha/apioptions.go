package alpha

type CpiInterval int64

const (
	Monthly CpiInterval = iota
	SemiAnnual
)

func (c CpiInterval) String() string {
	switch c {
	case Monthly:
		return "monthly"
	case SemiAnnual:
		return "semiannual"
	}
	return "unknown"
}
