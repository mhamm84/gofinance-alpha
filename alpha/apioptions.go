package alpha

type Interval int8

const (
	Daily Interval = iota + 1
	Monthly
	Quarterly
	SemiAnnual
	Annual
)

func (c Interval) String() string {
	switch c {
	case Daily:
		return "daily"
	case Monthly:
		return "monthly"
	case Quarterly:
		return "quarterly"
	case SemiAnnual:
		return "semiannual"
	case Annual:
		return "annual"
	}
	return "unknown"
}

type Maturity int8

const (
	ThreeMonth Maturity = iota + 1
	TwoYear
	FiveYear
	SevenYear
	TenYear
	ThirtyYear
)

func (c Maturity) String() string {
	switch c {
	case ThreeMonth:
		return "3month"
	case TwoYear:
		return "2year"
	case FiveYear:
		return "5year"
	case SevenYear:
		return "7year"
	case TenYear:
		return "10year"
	case ThirtyYear:
		return "30year"
	}
	return "unknown"
}
