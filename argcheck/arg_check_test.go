package argcheck

import (
	"testing"
	"time"
)

func TestNotNil_panic(t *testing.T) {
	assertPanic(t, func() {
		NotNil(nil, "msg")
	})
}

func TestNotNil_dont_panic(t *testing.T) {
	NotNil(time.Now(), "msg")
}

func ExampleNotNil() {
	address := struct {
		Street string
		Number int
	}{
		"Nicaragua",
		2533,
	}
	NotNil(address, "Address must be provided.")
}

func TestNotNilf_panic(t *testing.T) {
	assertPanic(t, func() {
		NotNilf(nil, "msg %v", "msgArg")
	})
}

func TestNotNilf_dont_panic(t *testing.T) {
	NotNilf(time.Now(), "msg %v", "msgArg")
}

func TestIsTrue_panic(t *testing.T) {
	assertPanic(t, func() {
		IsTrue(false, "msg")
	})
}

func TestIsTrue_dont_panic(t *testing.T) {
	IsTrue(true, "msg")
}

func TestIsTruef_panic(t *testing.T) {
	assertPanic(t, func() {
		IsTruef(false, "msg %v", "msgArg")
	})
}

func TestIsTruef_dont_panic(t *testing.T) {
	IsTruef(true, "msg %v", "msgArg")
}

func TestGreaterThan_panic(t *testing.T) {
	assertPanic(t, func() {
		GreaterThan(10, 100, "msg")
	})
}

func TestGreaterThan_dont_panic(t *testing.T) {
	GreaterThan(100, 10, "msg")
}

func TestGreaterThanf_panic(t *testing.T) {
	assertPanic(t, func() {
		GreaterThanf(10, 100, "msg %v", "msgArg")
	})
}

func TestGreaterThanf_dont_panic(t *testing.T) {
	GreaterThanf(100, 10, "msg %v", "msgArg")
}

func TestLessThan_panic(t *testing.T) {
	assertPanic(t, func() {
		LessThan(100, 10, "msg")
	})
}

func TestLessThan_dont_panic(t *testing.T) {
	LessThan(10, 100, "msg")
}

func TestLessThanf_panic(t *testing.T) {
	assertPanic(t, func() {
		LessThanf(100, 10, "msg %v", "msgArg")
	})
}

func TestLessThanf_dont_panic(t *testing.T) {
	LessThanf(10, 100, "msg %v", "msgArg")
}

func TestBetween_panic(t *testing.T) {
	assertPanic(t, func() {
		Between(1, 10, 100, "msg")
	})
}

func TestBetween_dont_panic(t *testing.T) {
	Between(50, 10, 100, "msg")
}

func TestBetweenf_panic(t *testing.T) {
	assertPanic(t, func() {
		Betweenf(1, 10, 100, "msg %v", "msgArg")
	})
}

func TestBetweenf_dont_panic(t *testing.T) {
	Betweenf(50, 10, 100, "msg %v", "msgArg")
}

func TestMatches_panic(t *testing.T) {
	assertPanic(t, func() {
		Matches("abc", "[d..z]", "msg")
	})
}

func TestMatches_dont_panic(t *testing.T) {
	Matches("abc", "[a..z]", "msg")
}

func TestMatchesf_panic(t *testing.T) {
	assertPanic(t, func() {
		Matchesf("abc", "[d..z]", "msg %v", "msgArg")
	})
}

func TestMatchesf_dont_panic(t *testing.T) {
	Matchesf("abc", "[a..z]", "msg %v", "msgArg")
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic expected")
		}
	}()
	f()
}
