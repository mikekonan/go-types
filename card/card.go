package card

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	layout                = "01/06"
	dateSeparator         = "/"
	maxCardDateExpireYear = 2099
	gatewayCentury        = 21
)

//CardDate represents as expired card date type with format MM/YY
type CardDate string

//Value implementation of driver.Valuer
func (cardDate CardDate) Value() (value driver.Value, err error) {
	if err = cardDate.Validate(); err != nil {
		return nil, err
	}

	return cardDate.String(), nil
}

//UnmarshalJSON unmarshall implementation for CardDate
func (cardDate *CardDate) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	var indexSplit = strings.Index(str, dateSeparator)
	if indexSplit != 2 || len(str) != 5 {
		return fmt.Errorf("invalid CardDate format. Format: MM/YY")
	}

	var _, err = time.Parse(layout, str)
	if err != nil {
		return fmt.Errorf("invalid CardDate: %w", err)
	}

	*cardDate = CardDate(str)

	return nil
}

//Validate implementation of ozzo-validation Validate interface
func (cardDate CardDate) Validate() error {
	if cardDate == "" {
		return fmt.Errorf("invalid CardDate: cannot be blank")
	}

	var realYear, realMonth, err = cardDate.parseDate()
	if err != nil {
		return fmt.Errorf("ivalid CardDate: %w", err)
	}

	var leftYear, leftMonth, _ = time.Now().Date()
	if leftYear > realYear || (leftYear == realYear && leftMonth > realMonth) {
		return fmt.Errorf("ivalid CardDate: card expired")
	}
	if realYear < leftYear || realYear > maxCardDateExpireYear {
		return fmt.Errorf("invalid CardDate format: year is not valid")
	}

	return nil
}

//String implementation of Stringer interface
func (cardDate CardDate) String() string {
	return string(cardDate)
}

func (cardDate CardDate) parseDate() (year int, month time.Month, err error) {
	if len(cardDate) != 5 {
		return 0, 0, fmt.Errorf("invalid CardDate format. Format: MM/YY")
	}
	separated := strings.Split(cardDate.String(), dateSeparator)
	if len(separated) != 2 {
		return 0, 0, fmt.Errorf("invalid CardDate format. Format: MM/YY")
	}
	parsedMonth := separated[0]
	parsedYear := separated[1]

	intMonth, err := strconv.Atoi(parsedMonth)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid CardDate format: month is not valid")
	}
	if intMonth > 12 || intMonth < 1 {
		return 0, 0, fmt.Errorf("invalid CardDate format: month is not valid")
	}

	intYear, err := strconv.Atoi(fmt.Sprintf("%d%s", gatewayCentury-1, parsedYear))
	if err != nil {
		return 0, 0, fmt.Errorf("invalid CardDate format: year is not valid")
	}

	return intYear, time.Month(intMonth), nil
}
