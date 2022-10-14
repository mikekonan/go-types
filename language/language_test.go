package language

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestLanguageByNameIsSet(t *testing.T) {
	for key, language := range LanguageByName {
		if !Name(key).IsSet() {
			t.FailNow()
		}

		if !language.name.IsSet() {
			t.FailNow()
		}

		if !language.alpha2.IsSet() {
			t.FailNow()
		}

		if !language.alpha3.IsSet() {
			t.FailNow()
		}
	}
}

func TestLanguageByAlpha3IsSet(t *testing.T) {
	for key, language := range LanguageByAlpha3 {
		if !Alpha3Code(key).IsSet() {
			t.FailNow()
		}

		if !language.name.IsSet() {
			t.FailNow()
		}

		if !language.alpha2.IsSet() {
			t.FailNow()
		}

		if !language.alpha3.IsSet() {
			t.FailNow()
		}
	}
}

func TestLanguageByAlpha2IsSet(t *testing.T) {
	for key, language := range LanguageByAlpha2 {
		if !Alpha2Code(key).IsSet() {
			t.FailNow()
		}

		if !language.name.IsSet() {
			t.FailNow()
		}

		if !language.alpha2.IsSet() {
			t.FailNow()
		}

		if !language.alpha3.IsSet() {
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
	for key, language := range LanguageByName {
		if Name(key) != language.name {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha3 {
		if Alpha3Code(key) != language.alpha3 {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha2 {
		if Alpha2Code(key) != language.alpha2 {
			t.FailNow()
		}
	}
}

func TestMappingStringsCorrect(t *testing.T) {
	for key, language := range LanguageByName {
		if Name(key).String() != language.name.String() {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha3 {
		if Alpha3Code(key).String() != language.alpha3.String() {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha2 {
		if Alpha2Code(key).String() != language.alpha2.String() {
			t.FailNow()
		}
	}
}

func TestMappingValueCorrect(t *testing.T) {
	for key, language := range LanguageByName {
		_, actual := Name(key).Value()
		_, expected := language.name.Value()

		if actual != expected {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha3 {
		_, actual := Alpha3Code(key).Value()
		_, expected := language.name.Value()

		if actual != expected {
			t.FailNow()
		}
	}

	for key, language := range LanguageByAlpha2 {
		_, actual := Alpha2Code(key).Value()
		_, expected := language.name.Value()

		if actual != expected {
			t.FailNow()
		}
	}
}

func TestWrongNameValue(t *testing.T) {
	value, err := Name("l").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestWrongAlpha2Value(t *testing.T) {
	value, err := Alpha2Code("l").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestWrongAlpha3Value(t *testing.T) {
	value, err := Alpha3Code("l").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestNameValidate(t *testing.T) {
	if Name("l").Validate() == nil {
		t.FailNow()
	}

	if Name("Latvian").Validate() != nil {
		t.FailNow()
	}
}

func TestAlpha2Validate(t *testing.T) {
	if Alpha2Code("l").Validate() == nil {
		t.FailNow()
	}

	if Alpha2Code("lv").Validate() != nil {
		t.FailNow()
	}
}

func TestAlpha3Validate(t *testing.T) {
	if Alpha3Code("lv").Validate() == nil {
		t.FailNow()
	}

	if Alpha3Code("lav").Validate() != nil {
		t.FailNow()
	}
}

func TestCountryTypes(t *testing.T) {
	for _, language := range LanguageByAlpha3 {
		if language.name != language.Name() {
			t.FailNow()
		}

		if language.alpha3 != language.Alpha3Code() {
			t.FailNow()
		}

		if language.alpha2 != language.Alpha2Code() {
			t.FailNow()
		}

		if string(language.name) != language.NameStr() {
			t.FailNow()
		}

		if string(language.alpha3) != language.Alpha3CodeStr() {
			t.FailNow()
		}

		if string(language.alpha2) != language.Alpha2CodeStr() {
			t.FailNow()
		}
	}
}

func TestAlpha2Lookup(t *testing.T) {
	if _, ok := ByAlpha2Code("l"); ok {
		t.FailNow()
	}

	if _, ok := ByAlpha2CodeStr("l"); ok {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeErr("l"); err == nil {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeStrErr("l"); err == nil {
		t.FailNow()
	}

	if _, ok := ByAlpha2Code(Latvian.Alpha2Code()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha2CodeStr(strings.ToUpper(Latvian.Alpha2Code().String())); !ok {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeErr(Latvian.Alpha2Code()); err != nil {
		t.FailNow()
	}

	if _, err := ByAlpha2CodeStrErr(strings.ToUpper(Latvian.Alpha2Code().String())); err != nil {
		t.FailNow()
	}
}

func TestAlpha3Lookup(t *testing.T) {
	if _, ok := ByAlpha3Code("l"); ok {
		t.FailNow()
	}

	if _, ok := ByAlpha3CodeStr("l"); ok {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeErr("l"); err == nil {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeStrErr("l"); err == nil {
		t.FailNow()
	}

	if _, ok := ByAlpha3Code(Latvian.Alpha3Code()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha3CodeStr(Latvian.Alpha3Code().String()); !ok {
		t.FailNow()
	}

	if _, ok := ByAlpha3CodeStr(strings.ToUpper(Latvian.Alpha3Code().String())); !ok {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeErr(Latvian.Alpha3Code()); err != nil {
		t.FailNow()
	}

	if _, err := ByAlpha3CodeStrErr(strings.ToUpper(Latvian.Alpha3Code().String())); err != nil {
		t.FailNow()
	}
}

func TestNameLookup(t *testing.T) {
	if _, ok := ByName("l"); ok {
		t.FailNow()
	}

	if _, ok := ByNameStr("l"); ok {
		t.FailNow()
	}

	if _, err := ByNameErr("l"); err == nil {
		t.FailNow()
	}

	if _, err := ByNameStrErr("l"); err == nil {
		t.FailNow()
	}

	if _, ok := ByName(Latvian.Name()); !ok {
		t.FailNow()
	}

	if _, ok := ByNameStr(Latvian.Name().String()); !ok {
		t.FailNow()
	}

	if _, err := ByNameErr(Latvian.Name()); err != nil {
		t.FailNow()
	}

	if _, err := ByNameStrErr(Latvian.Name().String()); err != nil {
		t.FailNow()
	}
}

func TestDriverValue(t *testing.T) {
	if value, err := Alpha2Code("lv").Value(); err != nil || value.(string) != Alpha2Code("LV").String() {
		t.FailNow()
	}

	if value, err := Alpha2Code("invalid").Value(); err == nil || value != nil {
		t.FailNow()
	}

	if value, err := Alpha2Code("").Value(); err != nil || value.(string) != "" {
		t.FailNow()
	}

	if value, err := Alpha3Code("lav").Value(); err != nil || value.(string) != Alpha3Code("LAV").String() {
		t.FailNow()
	}

	if value, err := Alpha3Code("invalid").Value(); err == nil || value != nil {
		t.FailNow()
	}

	if value, err := Alpha3Code("").Value(); err != nil || value.(string) != "" {
		t.FailNow()
	}

	if value, err := Name("Latvian").Value(); err != nil || value.(string) != Name("Latvian").String() {
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
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"name":"%s"}`, Latvian.Name())), &positive); err != nil && positive.Name != Latvian.Name() {
		t.FailNow()
	}

	var negative NameStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"name":"%s"}`, "RUsski")), &negative); err == nil {
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
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, Latvian.Alpha2Code())), &positive); err != nil && positive.Alpha2Code != Latvian.Alpha2Code() {
		t.FailNow()
	}

	var negative Alpha2CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "RY")), &negative); err == nil {
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
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, Latvian.Alpha3Code())), &positive); err != nil && positive.Alpha3Code != Latvian.Alpha3Code() {
		t.FailNow()
	}

	var negative Alpha3CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "RUZ")), &negative); err == nil {
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
