package timezone

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Timezone string

//UnmarshalJSON unmarshall implementation for timezone
func (name *Timezone) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	if _, err := ByNameStrErr(str); err != nil {
		return err
	}

	*name = Timezone(str)
	return nil
}

//Value implementation of driver.Valuer
func (name Timezone) Value() (value driver.Value, err error) {
	if name == "" {
		return "", nil
	}

	if _, err = ByNameStrErr(string(name)); err != nil {
		return nil, err
	}

	return name.String(), nil
}

//Validate implementation of ozzo-validation Validate interface
func (name Timezone) Validate() (err error) {
	_, err = ByNameStrErr(string(name))

	return
}

//String implementation of Stringer interface
func (name Timezone) String() string {
	return string(name)
}

//ByNameStr lookup for timezone by name string
func ByNameStr(timezoneStr string) (result Timezone, ok bool) {
	result, ok = timezonesByName[timezoneStr]
	return
}

//ByNameStrErr lookup for timezone by name with error return type
func ByNameStrErr(timezoneStr string) (result Timezone, err error) {
	var ok bool
	result, ok = ByNameStr(timezoneStr)
	if !ok {
		err = fmt.Errorf("'%s' is not valid RFC-6557 Timezone", timezoneStr)
	}

	return
}
