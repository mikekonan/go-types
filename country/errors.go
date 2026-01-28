package country

const (
	standardISO3166alpha2  = "ISO-3166-alpha2 code"
	standardISO3166alpha3  = "ISO-3166-alpha3 code"
	standardISO3166Country = "ISO-3166 Country name"
)

type InvalidDataError struct {
	data     string
	standard string
}

func (e InvalidDataError) Error() string {
	return "'" + string(e.data) + "' is not valid " + e.standard
}
