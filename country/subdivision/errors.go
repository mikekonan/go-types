package subdivision

import "strings"

const standardISO31662Code = "ISO-3166-2 subdivision code"

type InvalidDataError struct {
	data     string
	standard string
}

func (e InvalidDataError) Error() string {
	return "'" + string(e.data) + "' is not valid " + e.standard
}

// newInvalidDataError creates an InvalidDataError for the provided data and standard.
// The returned error's data field is a cloned copy of the input string to avoid retaining references to the original.
func newInvalidDataError(data string, standard string) InvalidDataError {
	// data copy to avoid retaining references to potentially unsafe strings
	return InvalidDataError{data: strings.Clone(data), standard: standard}
}
