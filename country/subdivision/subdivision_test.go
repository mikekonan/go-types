package subdivision

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func TestSubdivisionByCodeIsSet(t *testing.T) {
	for key, sub := range SubdivisionByCode {
		if !Code(key).IsSet() {
			t.FailNow()
		}

		if !sub.name.IsSet() {
			t.FailNow()
		}

		if !sub.code.IsSet() {
			t.FailNow()
		}

		if !sub.countryCode.IsSet() {
			t.FailNow()
		}

		if !sub.category.IsSet() {
			t.FailNow()
		}
	}
}

func TestIsNotSet(t *testing.T) {
	if Code("").IsSet() {
		t.FailNow()
	}

	if Name("").IsSet() {
		t.FailNow()
	}

	if Category("").IsSet() {
		t.FailNow()
	}
}

func TestMappingIsCorrect(t *testing.T) {
	for key, sub := range SubdivisionByCode {
		if Code(key) != sub.code {
			t.FailNow()
		}
	}
}

func TestMappingStringsCorrect(t *testing.T) {
	for key, sub := range SubdivisionByCode {
		if Code(key).String() != sub.code.String() {
			t.FailNow()
		}
	}
}

func TestMappingValueCorrect(t *testing.T) {
	for key, sub := range SubdivisionByCode {
		actualValue, actualErr := Code(key).Value()
		expectedValue, expectedErr := sub.code.Value()

		if actualErr != nil {
			t.Fatalf("unexpected error for Code(%q).Value(): %v", key, actualErr)
		}
		if expectedErr != nil {
			t.Fatalf("unexpected error for sub.code.Value() (%q): %v", sub.code, expectedErr)
		}
		if actualValue != expectedValue {
			t.Fatalf("value mismatch for %q: got %v, want %v", key, actualValue, expectedValue)
		}
	}
}

func TestWrongCodeValue(t *testing.T) {
	value, err := Code("a").Value()
	if err == nil || value != nil {
		t.FailNow()
	}
}

func TestCodeValidate(t *testing.T) {
	if Code("a").Validate() == nil {
		t.FailNow()
	}

	if Code("US-CA").Validate() != nil {
		t.FailNow()
	}
}

func TestSubdivisionTypes(t *testing.T) {
	for _, sub := range SubdivisionByCode {
		if sub.name != sub.Name() {
			t.FailNow()
		}

		if sub.code != sub.Code() {
			t.FailNow()
		}

		if sub.countryCode != sub.CountryCode() {
			t.FailNow()
		}

		if sub.category != sub.Category() {
			t.FailNow()
		}

		if string(sub.name) != sub.NameStr() {
			t.FailNow()
		}

		if string(sub.code) != sub.CodeStr() {
			t.FailNow()
		}

		if string(sub.countryCode) != sub.CountryCodeStr() {
			t.FailNow()
		}

		if string(sub.category) != sub.CategoryStr() {
			t.FailNow()
		}
	}
}

func TestCodeLookup(t *testing.T) {
	if _, ok := ByCode("a"); ok {
		t.FailNow()
	}

	if _, ok := ByCodeStr("a"); ok {
		t.FailNow()
	}

	if _, err := ByCodeErr("a"); err == nil {
		t.FailNow()
	}

	if _, err := ByCodeStrErr("a"); err == nil {
		t.FailNow()
	}

	if _, ok := ByCode(USCA.Code()); !ok {
		t.FailNow()
	}

	if _, ok := ByCodeStr(strings.ToLower(USCA.Code().String())); !ok {
		t.FailNow()
	}

	if _, err := ByCodeErr(USCA.Code()); err != nil {
		t.FailNow()
	}

	if _, err := ByCodeStrErr(strings.ToLower(USCA.Code().String())); err != nil {
		t.FailNow()
	}
}

func TestCountryCodeLookup(t *testing.T) {
	if _, ok := ByCountryCodeStr("ZZ"); ok {
		t.FailNow()
	}

	if _, err := ByCountryCodeStrErr("ZZ"); err == nil {
		t.FailNow()
	}

	subs, ok := ByCountryCodeStr("US")
	if !ok {
		t.FailNow()
	}

	if len(subs) < 50 {
		t.Errorf("expected at least 50 US subdivisions, got %d", len(subs))
	}

	subs, err := ByCountryCodeStrErr("CA")
	if err != nil {
		t.FailNow()
	}

	if len(subs) < 10 {
		t.Errorf("expected at least 10 CA subdivisions, got %d", len(subs))
	}
}

func TestDriverValue(t *testing.T) {
	if value, err := Code("US-CA").Value(); err != nil || value.(string) != Code("US-CA").String() {
		t.FailNow()
	}

	if value, err := Code("invalid").Value(); err == nil || value != nil {
		t.FailNow()
	}

	if value, err := Code("").Value(); err != nil || value.(string) != "" {
		t.FailNow()
	}
}

func TestCodeUnmarshalJson(t *testing.T) {
	type CodeStruct struct {
		Code Code `json:"code"`
	}

	var positive CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, USCA.Code())), &positive); err != nil || positive.Code != USCA.Code() {
		t.FailNow()
	}

	var negative CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "XX-YY")), &negative); err == nil {
		t.FailNow()
	}

	var lowercase CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "us-ca")), &lowercase); err != nil || lowercase.Code != USCA.Code() {
		t.FailNow()
	}

	var wrongCode CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "wrong code")), &wrongCode); err == nil {
		t.FailNow()
	}

	var emptyCode CodeStruct
	if err := json.Unmarshal([]byte(fmt.Sprintf(`{"code":"%s"}`, "")), &emptyCode); err != nil {
		t.FailNow()
	}
}

func TestNameUnmarshalJson(t *testing.T) {
	type NameStruct struct {
		Name Name `json:"name"`
	}

	var positive NameStruct
	if err := json.Unmarshal([]byte(`{"name":"California"}`), &positive); err != nil || positive.Name != "California" {
		t.FailNow()
	}

	var emptyName NameStruct
	if err := json.Unmarshal([]byte(`{"name":""}`), &emptyName); err != nil {
		t.FailNow()
	}
}

func TestCategoryUnmarshalJson(t *testing.T) {
	type CategoryStruct struct {
		Category Category `json:"category"`
	}

	var positive CategoryStruct
	if err := json.Unmarshal([]byte(`{"category":"state"}`), &positive); err != nil || positive.Category != "state" {
		t.FailNow()
	}

	var empty CategoryStruct
	if err := json.Unmarshal([]byte(`{"category":""}`), &empty); err != nil {
		t.FailNow()
	}
}

func TestSubdivisionsByCountryCount(t *testing.T) {
	// Spot-check known counts
	us, ok := ByCountryCodeStr("US")
	if !ok {
		t.Fatal("US not found")
	}
	if len(us) < 50 {
		t.Errorf("expected US to have at least 50 subdivisions, got %d", len(us))
	}

	ca, ok := ByCountryCodeStr("CA")
	if !ok {
		t.Fatal("CA not found")
	}
	if len(ca) != 13 {
		t.Errorf("expected CA to have 13 subdivisions, got %d", len(ca))
	}

	de, ok := ByCountryCodeStr("DE")
	if !ok {
		t.Fatal("DE not found")
	}
	if len(de) != 16 {
		t.Errorf("expected DE to have 16 subdivisions, got %d", len(de))
	}
}

func TestValidateForCountry(t *testing.T) {
	// valid code for correct country
	if err := Code("US-CA").ValidateForCountry("US"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// valid code for wrong country
	if err := Code("US-CA").ValidateForCountry("DE"); err == nil {
		t.Fatal("expected error for US-CA with country DE")
	}

	// invalid code
	if err := Code("XX-YY").ValidateForCountry("US"); err == nil {
		t.Fatal("expected error for invalid code XX-YY")
	}

	// case insensitive
	if err := Code("us-ca").ValidateForCountry("US"); err != nil {
		t.Fatalf("expected no error for lowercase us-ca, got %v", err)
	}

	// all subdivisions match their own country
	for _, sub := range SubdivisionByCode {
		if err := sub.code.ValidateForCountry(sub.countryCode); err != nil {
			t.Fatalf("subdivision %s should be valid for country %s: %v", sub.code, sub.countryCode, err)
		}
	}
}

func TestTotalSubdivisionCount(t *testing.T) {
	if len(SubdivisionByCode) < 4000 {
		t.Errorf("expected at least 4000 subdivisions, got %d", len(SubdivisionByCode))
	}
}
