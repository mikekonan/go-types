package subdivision

// InvalidDataError represents an error when invalid data is provided for a given standard.
type InvalidDataError struct {
	data     string
	standard string
}

func (e InvalidDataError) Error() string {
	return "'" + e.data + "' is not valid " + e.standard
}
