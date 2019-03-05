package require

import (
	"reflect"
	"testing"
)

func NoError(t *testing.T, err error, msgs ...string) {
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
		for _, msg := range msgs {
			t.Errorf("\n" + msg)
		}
		t.FailNow()
	}
}

func NotNil(t *testing.T, object interface{}, msgs ...string) {
	if isNil(object) {
		t.Errorf("Expected nil, but got an object")
		for _, msg := range msgs {
			t.Errorf("\n" + msg)
		}
		t.FailNow()
	}
}

func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

func EqualError(t *testing.T, theError error, errString string, msgs ...string) {
	if theError == nil {
		t.Errorf("Expected an error but got nil")
		for _, msg := range msgs {
			t.Errorf("\n" + msg)
		}
		t.FailNow()
		return
	}

	if theError.Error() != errString {
		t.Errorf("Expected an error with message:\n%qbut got \n%q", errString, theError)
		for _, msg := range msgs {
			t.Errorf("\n" + msg)
		}
		t.FailNow()
	}
}

func EqualFloat64(t *testing.T, expected, actual float64, msgs ...string) {
	if expected != actual {
		t.Errorf("Expected %f but got %f", expected, actual)
		for _, msg := range msgs {
			t.Errorf("\n" + msg)
		}
		t.FailNow()
	}
}

func EqualStringPtr(t *testing.T, expected, actual *string, msgs ...string) {
	if expected == nil && actual == nil {
		return
	}
	if expected != nil && actual != nil {
		EqualString(t, *expected, *actual, msgs...)
		return
	}

	t.Errorf("Expected %v but got %v", expected, actual)
	for _, msg := range msgs {
		t.Errorf("\n" + msg)
	}
	t.FailNow()
}

func EqualString(t *testing.T, expected, actual string, msgs ...string) {
	if expected != actual {
		t.Errorf("Expected %s but got %s", expected, actual)
		for _, msg := range msgs {
			t.Errorf("\n" + msg)
		}
		t.FailNow()
	}
}

func EqualUInt32(t *testing.T, expected, actual uint32, msgs ...string) {
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		for _, msg := range msgs {
			t.Errorf("\n" + msg)
		}
		t.FailNow()
	}
}

func EqualUInt64Ptr(t *testing.T, expected, actual *uint64, msgs ...string) {
	if expected == nil && actual == nil {
		return
	}
	if expected != nil && actual != nil {
		EqualUInt64(t, *expected, *actual, msgs...)
		return
	}

	t.Errorf("Expected %v but got %v", expected, actual)
	for _, msg := range msgs {
		t.Errorf("\n" + msg)
	}
	t.FailNow()
}

func EqualUInt64(t *testing.T, expected, actual uint64, msgs ...string) {
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		for _, msg := range msgs {
			t.Errorf("\n" + msg)
		}
		t.FailNow()
	}
}

func EqualIntPtr(t *testing.T, expected, actual *int, msgs ...string) {
	if expected == nil && actual == nil {
		return
	}
	if expected != nil && actual != nil {
		EqualInt(t, *expected, *actual, msgs...)
		return
	}

	t.Errorf("Expected %v but got %v", expected, actual)
	for _, msg := range msgs {
		t.Errorf("\n" + msg)
	}
	t.FailNow()
}

func EqualInt(t *testing.T, expected, actual int, msgs ...string) {
	if expected != actual {
		t.Errorf("Expected %d but got %d", expected, actual)
		for _, msg := range msgs {
			t.Errorf("\n" + msg)
		}
		t.FailNow()
	}
}
