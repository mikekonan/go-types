package currency

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Currency represents a currency type from ISO-4217
type Currency string

// Value implementation of driver.Valuer
func (currency Currency) Value() (value driver.Value, err error) {
	if currency == "" {
		return "", nil
	}

	if err = currency.Validate(); err != nil {
		return nil, err
	}

	return currency.String(), nil
}

// UnmarshalJSON unmarshall implementation for Currency
func (currency *Currency) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	currencyValue, err := ByCurrencyStrErr(str)
	if err != nil {
		return err
	}

	*currency = currencyValue.Currency()

	return nil
}

// Validate implementation of ozzo-validation Validate interface
func (currency Currency) Validate() error {
	if _, ok := ByCurrencyStr(string(currency)); !ok {
		return fmt.Errorf("'%s' is not valid ISO-4217 currency", currency)
	}

	return nil
}

// IsSet indicates if Currency is set
func (currency Currency) IsSet() bool {
	return len(string(currency)) > 0
}

// String implementation of Stringer interface
func (currency Currency) String() string {
	return string(currency)
}
