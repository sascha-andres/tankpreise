package tankpreise

import "strings"

// String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
func (pr *PricesRequest) String() string {
	return strings.Join(pr.IDs, ",")
}
