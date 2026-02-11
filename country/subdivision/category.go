package subdivision

import (
	"database/sql/driver"
	"fmt"

	"github.com/mikekonan/go-types/v2/internal/utils"
)

// Category represents subdivision category (e.g. "state", "province", "territory")
type Category string

// UnmarshalJSON unmarshall implementation for category
func (category *Category) UnmarshalJSON(data []byte) error {
	str, isEmptyValue, err := utils.UnsafeStringFromJson(data)
	if err != nil {
		return err
	}

	if isEmptyValue {
		return nil
	}

	if _, ok := subdivisionCategories[str]; !ok {
		return fmt.Errorf("'%s' is not valid ISO-3166-2 subdivision category", str)
	}

	*category = Category(str)
	return nil
}

// Validate implementation of ozzo-validation Validate interface
func (category Category) Validate() error {
	if _, ok := subdivisionCategories[string(category)]; !ok {
		return fmt.Errorf("'%s' is not valid ISO-3166-2 subdivision category", category)
	}

	return nil
}

// Value implementation of driver.Valuer
func (category Category) Value() (value driver.Value, err error) {
	if category == "" {
		return "", nil
	}

	if err = category.Validate(); err != nil {
		return nil, err
	}

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
