package time

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"

	"gopkg.in/guregu/null.v4"
)

type Time null.Time

func NullTime(t time.Time) Time {
	return Time(null.TimeFrom(t))
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	return t.parseTimeFromString(str)
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.format())
}

func (t *Time) UnmarshalXML(decoder *xml.Decoder, element xml.StartElement) error {
	var str string
	err := decoder.DecodeElement(&str, &element)
	if err != nil {
		return err
	}

	return t.parseTimeFromString(str)
}

func (t Time) MarshalXML(decoder *xml.Encoder, element xml.StartElement) error {
	return decoder.EncodeElement(t.format(), element)
}

func (t Time) String() string {
	var str = t.format()
	if str != nil {
		return *str
	}
	return ""
}

func (t Time) format() *string {
	if t.Valid {
		var str = t.Time.Format(time.RFC3339Nano)
		return &str
	}

	return nil
}

func (t *Time) parseTimeFromString(value string) (err error) {
	if value == "" {
		return nil
	}

	var newTime time.Time
	if newTime, err = time.Parse(time.RFC3339Nano, value); err != nil {
		if errTime, ok := err.(*time.ParseError); ok {
			return fmt.Errorf("cannot parse time '%s' invalid format '%s'", errTime.Value, errTime.Layout)
		}
		return err
	}

	*t = NullTime(newTime)
	return nil
}

// Validate implementation of ozzo-validation Validate interface
func (t Time) Validate() (err error) {
	if !t.Valid {
		return fmt.Errorf("'%s' is not valid time", t)
	}

	return
}

func ParseNullTimeFromString(value string) (t Time, err error) {
	if err = t.parseTimeFromString(value); err != nil {
		return t, err
	}
	return t, nil
}
