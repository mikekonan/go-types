package country

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestCountryByNameIsSet(t *testing.T) {
	for key, country := range CountryByName {
		if !Name(key).IsSet() {
			t.FailNow()
		}

		if !country.name.IsSet() {
			t.FailNow()
		}

		if !country.alpha2.IsSet() {
			t.FailNow()
		}

		if !country.alpha3.IsSet() {
			t.FailNow()
		}
	}
}

func TestCountryByAlpha3IsSet(t *testing.T) {
	for key, country := range CountryByAlpha3 {
		if !Alpha3Code(key).IsSet() {
			t.FailNow()
		}

		if !country.name.IsSet() {
			t.FailNow()
		}

		if !country.alpha2.IsSet() {
			t.FailNow()
		}

		if !country.alpha3.IsSet() {
			t.FailNow()
		}
	}
}

func TestCountryByAlpha2IsSet(t *testing.T) {
	for key, country := range CountryByAlpha2 {
		if !Alpha2Code(key).IsSet() {
			t.FailNow()
		}

		if !country.name.IsSet() {
			t.FailNow()
		}

		if !country.alpha2.IsSet() {
			t.FailNow()
		}

		if !country.alpha3.IsSet() {
			t.FailNow()
		}
	}
}

func TestIsNotSet(t *testing.T) {
	if Name("").IsSet() {
		t.FailNow()
	}

	if Alpha2Code("").IsSet() {
		t.FailNow()
	}

	if Alpha3Code("").IsSet() {
		t.FailNow()
	}
}

func TestMappingIsCorrect(t *testing.T) {
	for key, country := range CountryByName {
		if Name(key) != country.name {
			t.FailNow()
		}
	}

	for key, country := range CountryByAlpha3 {
		if Alpha3Code(key) != country.alpha3 {
			t.FailNow()
		}
	}

	for key, country := range CountryByAlpha2 {
		if Alpha2Code(key) != country.alpha2 {
			t.FailNow()
		}
	}
}

func TestMappingStringsCorrect(t *testing.T) {
	for key, country := range CountryByName {
		if Name(key).String() != country.name.String() {
			t.FailNow()
		}
	}

	for key, country := range CountryByAlpha3 {
		if Alpha3Code(key).String() != country.alpha3.String() {
			t.FailNow()
		}
	}

	for key, country := range CountryByAlpha2 {
		if Alpha2Code(key).String() != country.alpha2.String() {
			t.FailNow()
		}
	}
}

func TestMappingValueCorrect(t *testing.T) {
	for key, country := range CountryByName {
		_, actual := Name(key).Value()
		_, expected := country.name.Value()

		if actual != expected {
			t.FailNow()
		}
	}

	for key, country := range CountryByAlpha3 {
		_, actual := Alpha3Code(key).Value()
		_, expected := country.name.Value()

		if actual != expected {
			t.FailNow()
		}
	}

	for key, country := range CountryByAlpha2 {
		_, actual := Alpha2Code(key).Value()
		_, expected := country.name.Value()

		if actual != expected {
			t.FailNow()
		}
	}
}

func TestWrongNameValue(t *testing.T) {
	value, err := Name("a").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestWrongAlpha2Value(t *testing.T) {
	value, err := Alpha2Code("a").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestWrongAlpha3Value(t *testing.T) {
	value, err := Alpha3Code("a").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestNameValidate(t *testing.T) {
	if Name("a").Validate() == nil {
		t.FailNow()
	}

	if Name("Brazil").Validate() != nil {
		t.FailNow()
	}
}

func TestAlpha2Validate(t *testing.T) {
	if Alpha2Code("a").Validate() == nil {
		t.FailNow()
	}

	if Alpha2Code("CM").Validate() != nil {
		t.FailNow()
	}
}

func TestAlpha3Validate(t *testing.T) {
	if Alpha3Code("a").Validate() == nil {
		t.FailNow()
	}

	if Alpha3Code("ARG").Validate() != nil {
		t.FailNow()
	}
}

func TestCountryTypes(t *testing.T) {
	for _, country := range CountryByAlpha3 {
		if country.name != country.Name() {
			t.FailNow()
		}

		if country.alpha3 != country.Alpha3Code() {
			t.FailNow()
		}

		if country.alpha2 != country.Alpha2Code() {
			t.FailNow()
		}

		if string(country.name) != country.NameStr() {
			t.FailNow()
		}

		if string(country.alpha3) != country.Alpha3CodeStr() {
			t.FailNow()
		}

		if string(country.alpha2) != country.Alpha2CodeStr() {
			t.FailNow()
		}
	}
}

func TestAlpha2Lookup(t *testing.T) {
	if _, ok := ByAlpha2Code("a"); ok {
		t.FailNow()
	}

	if _, ok := ByAlpha2CodeStr("a"); ok {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeErr("a"); err == nil {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeStrErr("a"); err == nil {
		t.FailNow()
	}

	if _, ok := ByAlpha2Code(Canada.Alpha2Code()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha2CodeStr(strings.ToLower(Canada.Alpha2Code().String())); !ok {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeErr(Canada.Alpha2Code()); err != nil {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeStrErr(strings.ToLower(Canada.Alpha2Code().String())); err != nil {
		t.FailNow()
	}
}

func TestAlpha3Lookup(t *testing.T) {
	if _, ok := ByAlpha3Code("a"); ok {
		t.FailNow()
	}

	if _, ok := ByAlpha3CodeStr("a"); ok {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeErr("a"); err == nil {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeStrErr("a"); err == nil {
		t.FailNow()
	}

	if _, ok := ByAlpha3Code(Canada.Alpha3Code()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha3CodeStr(Canada.Alpha3Code().String()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha3CodeStr(strings.ToLower(Canada.Alpha3Code().String())); !ok {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeErr(Canada.Alpha3Code()); err != nil {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeStrErr(strings.ToLower(Canada.Alpha3Code().String())); err != nil {
		t.FailNow()
	}
}

func TestNameLookup(t *testing.T) {
	if _, ok := ByName("a"); ok {
		t.FailNow()
	}

	if _, ok := ByNameStr("a"); ok {
		t.FailNow()
	}

	if _, err := ByNameErr("a"); err == nil {
		t.FailNow()
	}

	if _, err := ByNameStrErr("a"); err == nil {
		t.FailNow()
	}

	if _, ok := ByName(Canada.Name()); !ok {
		t.FailNow()
	}

	if _, ok := ByNameStr(Canada.Name().String()); !ok {
		t.FailNow()
	}

	if _, err := ByNameErr(Canada.Name()); err != nil {
		t.FailNow()
	}

	if _, err := ByNameStrErr(Canada.Name().String()); err != nil {
		t.FailNow()
	}
}

func TestDriverValue(t *testing.T) {
	if value, err := Alpha2Code("US").Value(); err != nil || value.(string) != Alpha2Code("US").String() {
		t.FailNow()
	}

	if value, err := Alpha2Code("invalid").Value(); err == nil || value != nil {
		t.FailNow()
	}

	if value, err := Alpha2Code("").Value(); err != nil || value.(string) != "" {
		t.FailNow()
	}

	if value, err := Alpha3Code("USA").Value(); err != nil || value.(string) != Alpha3Code("USA").String() {
		t.FailNow()
	}

	if value, err := Alpha3Code("invalid").Value(); err == nil || value != nil {
		t.FailNow()
	}

	if value, err := Alpha3Code("").Value(); err != nil || value.(string) != "" {
		t.FailNow()
	}

	if value, err := Name("United States of America (the)").Value(); err != nil || value.(string) != Name("United States of America (the)").String() {
		t.FailNow()
	}

	if value, err := Name("invalid").Value(); err == nil || value != nil {
		t.FailNow()
	}

	if value, err := Name("").Value(); err != nil || value.(string) != "" {
		t.FailNow()
	}

}

func TestNameUnmarshalJson(t *testing.T) {
	type NameStruct struct {
		Name Name `json:"name"`
	}

	var positive NameStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"name":"%s"}`, UnitedStatesOfAmerica.Name())), &positive); err != nil && positive.Name != UnitedStatesOfAmerica.Name() {
		t.FailNow()
	}

	var negative NameStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"name":"%s"}`, "Raaa-siiii-yaaaa")), &negative); err == nil {
		t.FailNow()
	}

	var wrongName NameStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"name":"%s"}`, "wrong name")), &wrongName); err == nil {
		t.FailNow()
	}

	var emptyName NameStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"name":"%s"}`, "")), &emptyName); err != nil {
		t.FailNow()
	}
}

func TestAlpha2CodeUnmarshalJson(t *testing.T) {
	type Alpha2CodeStruct struct {
		Alpha2Code Alpha2Code `json:"code"`
	}

	var positive Alpha2CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, UnitedStatesOfAmerica.Alpha2Code())), &positive); err != nil && positive.Alpha2Code != UnitedStatesOfAmerica.Alpha2Code() {
		t.FailNow()
	}

	var negative Alpha2CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "RY")), &negative); err == nil {
		t.FailNow()
	}

	var lowercase Alpha2CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "ca")), &lowercase); err != nil || lowercase.Alpha2Code != Canada.Alpha2Code() {
		t.FailNow()
	}

	var wrongCode Alpha2CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "wrong code")), &wrongCode); err == nil {
		t.FailNow()
	}

	var emptyCode Alpha2CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "")), &emptyCode); err != nil {
		t.FailNow()
	}
}

func TestAlpha3CodeUnmarshalJson(t *testing.T) {
	type Alpha3CodeStruct struct {
		Alpha3Code Alpha3Code `json:"code"`
	}

	var positive Alpha3CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, UnitedStatesOfAmerica.Alpha3Code())), &positive); err != nil && positive.Alpha3Code != UnitedStatesOfAmerica.Alpha3Code() {
		t.FailNow()
	}

	var negative Alpha3CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "RUZ")), &negative); err == nil {
		t.FailNow()
	}

	var lowercase Alpha3CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "can")), &lowercase); err != nil || lowercase.Alpha3Code != Canada.Alpha3Code() {
		t.FailNow()
	}

	var wrongCode Alpha3CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "wrong code")), &wrongCode); err == nil {
		t.FailNow()
	}

	var emptyCode Alpha3CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "")), &emptyCode); err != nil {
		t.FailNow()
	}
}
