package timezone

import "testing"

func TestTimezone_MappingIsCorrect(t *testing.T) {
	for key, timezone := range timezonesByName {
		if key != string(timezone) {
			t.FailNow()
		}
	}
}

func TestTimezone_Value(t *testing.T) {
	value, err := Timezone("a").Value()
	if err == nil || value != nil {
		t.FailNow()
	}

	value, err = Timezone("Europe/Minsk").Value()
	if err != nil || value == nil {
		t.FailNow()
	}
}

func TestTimezone_Validate(t *testing.T) {
	if Timezone("a").Validate() == nil {
		t.FailNow()
	}

	if Timezone("Europe/Minsk").Validate() != nil {
		t.FailNow()
	}
}

func TestTimezone_Lookup(t *testing.T) {

	if _, ok := ByNameStr("a"); ok {
		t.FailNow()
	}

	if _, err := ByNameStrErr("a"); err == nil {
		t.FailNow()
	}

	if _, ok := ByNameStr("Europe/Minsk"); !ok {
		t.FailNow()
	}

	if _, err := ByNameStrErr("Europe/Minsk"); err != nil {
		t.FailNow()
	}
}
