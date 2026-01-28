package currency

import "strings"

const (
	standardISO4217Code     = "ISO-4217 code"
	standardISO4217Currency = "ISO-4217 currency"
	standardISO4217Number   = "ISO-4217 number"
	standardISO4217Country  = "ISO-4217 country"
)

type InvalidDataError struct {
	data     string
	standard string
}

func (e InvalidDataError) Error() string {
	return "'" + string(e.data) + "' is not valid " + e.standard
}

func newInvalidDataError(data string, standard string) InvalidDataError {
	// data copy to avoid retaining references to potentially unsafe strings
	return InvalidDataError{data: strings.Clone(data), standard: standard}
}
