package argcheck_test

import (
	"testing"
	"time"

	"github.com/raulbajales/go-argcheck/argcheck"
)

func TestNotNil_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.NotNil(nil, "msg")
	})
}

func TestNotNil_dont_panic(t *testing.T) {
	argcheck.NotNil(time.Now(), "msg")
}

func ExampleNotNil() {
	address := struct {
		Street string
		Number int
	}{
		"Nicaragua",
		2533,
	}
	argcheck.NotNil(address, "address must be provided")
}

func TestNotNilf_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.NotNilf(nil, "msg %v", "msgArg")
	})
}

func ExampleNotNilf() {
	moment := struct {
		t time.Time
	}{
		time.Now(),
	}
	argcheck.NotNilf(moment, "moment must be provided, like now: %v", time.Now())
}

func TestNotNilf_dont_panic(t *testing.T) {
	argcheck.NotNilf(time.Now(), "msg %v", "msgArg")
}

func TestIsTrue_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.IsTrue(false, "msg")
	})
}

func ExampleIsTrue() {
	countryCode := "AR"
	argcheck.IsTrue(countryCode != "GB" && countryCode != "BR", "GB and BR country codes are not allowed")
}

func TestIsTrue_dont_panic(t *testing.T) {
	argcheck.IsTrue(true, "msg")
}

func TestIsTruef_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.IsTruef(false, "msg %v", "msgArg")
	})
}

func ExampleIsTruef() {
	countryCode := "AR"
	argcheck.IsTruef(countryCode != "GB" && countryCode != "BR", "GB and BR country codes are not allowed, countryCode is %v", countryCode)
}

func TestIsTruef_dont_panic(t *testing.T) {
	argcheck.IsTruef(true, "msg %v", "msgArg")
}

func TestGreaterThan_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.GreaterThan(10, 100, "msg")
	})
}

func ExampleGreaterThan() {
	age := 12
	argcheck.GreaterThan(age, 16, "Cannot get driver licence")
}

func TestGreaterThan_dont_panic(t *testing.T) {
	argcheck.GreaterThan(100, 10, "msg")
}

func TestGreaterThanf_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.GreaterThanf(10, 100, "msg %v", "msgArg")
	})
}

func ExampleGreaterThanf() {
	age := 12
	argcheck.GreaterThanf(age, 16, "Cannot get driver licence, age is %v", age)
}

func TestGreaterThanf_dont_panic(t *testing.T) {
	argcheck.GreaterThanf(100, 10, "msg %v", "msgArg")
}

func TestLessThan_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.LessThan(100, 10, "msg")
	})
}

func ExampleLessThan() {
	age := 40
	argcheck.LessThan(age, 16, "TV show not recommended for adults")
}

func TestLessThan_dont_panic(t *testing.T) {
	argcheck.LessThan(10, 100, "msg")
}

func TestLessThanf_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.LessThanf(100, 10, "msg %v", "msgArg")
	})
}

func ExampleLessThanf() {
	age := 40
	argcheck.LessThanf(age, 16, "TV show not suitable for adults, age is %v", age)
}

func TestLessThanf_dont_panic(t *testing.T) {
	argcheck.LessThanf(10, 100, "msg %v", "msgArg")
}

func TestBetween_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.Between(1, 10, 100, "msg")
	})
}

func ExampleBetween() {
	minutesReadingPerDay := 16
	argcheck.Between(minutesReadingPerDay, 15, 120, "Only average reader allowed")
}

func TestBetween_dont_panic(t *testing.T) {
	argcheck.Between(50, 10, 100, "msg")
}

func TestBetweenf_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.Betweenf(1, 10, 100, "msg %v", "msgArg")
	})
}

func ExampleBetweenf() {
	minutesReadingPerDay := 16
	argcheck.Betweenf(minutesReadingPerDay, 15, 120, "Only average reader allowed, minutesReadingPerDay is %v", minutesReadingPerDay)
}

func TestBetweenf_dont_panic(t *testing.T) {
	argcheck.Betweenf(50, 10, 100, "msg %v", "msgArg")
}

func TestMatches_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.Matches("abc", "[d..z]", "msg")
	})
}

func ExampleMatches() {
	name := "Raul"
	argcheck.Matches(name, "[R].*", "Only names starting with R allowed")
}

func TestMatches_dont_panic(t *testing.T) {
	argcheck.Matches("abc", "[a..z]", "msg")
}

func TestMatchesf_panic(t *testing.T) {
	assertPanic(t, func() {
		argcheck.Matchesf("abc", "[d..z]", "msg %v", "msgArg")
	})
}

func ExampleMatchesf() {
	name := "Raul"
	argcheck.Matchesf(name, "[R].*", "Only names starting with R allowed, name is %v", name)
}

func TestMatchesf_dont_panic(t *testing.T) {
	argcheck.Matchesf("abc", "[a..z]", "msg %v", "msgArg")
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic expected")
		}
	}()
	f()
}
