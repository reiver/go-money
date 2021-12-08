package money

import (
	"database/sql/driver"
)

// Value is to make it so money.CAD fits the database/sql/driver.Valuer interface.
func (m CAD) Value() (driver.Value, error) {
	s := m.String()
	return s, nil
}
// Value is to make it so money.USD fits the database/sql/driver.Valuer interface.
func (m USD) Value() (driver.Value, error) {
	s := m.String()
	return s, nil
}
