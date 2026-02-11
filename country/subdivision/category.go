package subdivision

import (
	"database/sql/driver"
	"encoding/json"
)

// Category represents subdivision category (e.g. "state", "province", "territory")
type Category string

// UnmarshalJSON unmarshall implementation for category
func (category *Category) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	*category = Category(str)
	return nil
}

// Value implementation of driver.Valuer
func (category Category) Value() (value driver.Value, err error) {
	return category.String(), nil
}

// IsSet indicates if Category is set
func (category Category) IsSet() bool {
	return len(string(category)) > 0
}

// String implementation of Stringer interface
func (category Category) String() string {
	return string(category)
}
