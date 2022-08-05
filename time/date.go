package time

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"

	"gopkg.in/guregu/null.v4"
)

const format = "2006-01-02"

type Date null.Time

func DateFromTime(t time.Time) Date {
	return Date(null.TimeFrom(t))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	*d, err = ParseDateFromString(str)

	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.format())
}

func (d *Date) UnmarshalXML(decoder *xml.Decoder, element xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &element)
	if err != nil {
		return err
	}

	*d, err = ParseDateFromString(str)

	return err
}

func (d Date) MarshalXML(decoder *xml.Encoder, element xml.StartElement) error {
	return decoder.EncodeElement(d.format(), element)
}

func (d Date) String() string {
	var str = d.format()
	if str != nil {
		return *str
	}
	return ""
}

func (d Date) format() *string {
	if d.Valid {
		var str = d.Time.Format(format)
		return &str
	}

	return nil
}

// Validate implementation of ozzo-validation Validate interface
func (d Date) Validate() (err error) {
	if !d.Valid {
		return fmt.Errorf("'%s' is not valid date", d)
	}

	return
}

func ParseDateFromString(value string) (Date, error) {
	if value == "" {
		return Date{}, nil
	}

	var tm, err = time.Parse(format, value)
	if err != nil {
		return Date{}, err
	}

	return Date(null.TimeFrom(tm)), nil
}
