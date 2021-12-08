package money

import (
	"fmt"
)

// GoString returns a money.CAD in a Go code representation; ex: "money.USDCents(1234)".
func (m CAD) GoString() string {
	if nothingCAD() == m {
		return "money.mothingCAD()"
	}

	return fmt.Sprintf("money.CADCents(%d)", m.value)
}

// GoString returns a money.USD in a Go code representation; ex: "money.USDCents(1234)".
func (m USD) GoString() string {
	if nothingUSD() == m {
		return "money.nothingUSD()"
	}

	return fmt.Sprintf("money.USDCents(%d)", m.value)
}
