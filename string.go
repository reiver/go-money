package money


// String returns a string representing the CAD money in the form "$1.23".
func (m CAD) String() string {
	return m.Format("")
}


// String returns a string representing the USD money in the form "$1.23".
func (m USD) String() string {
	return m.Format("")
}
