package argcheck

import (
	"testing"
	"time"
)

func Test_NotNil_panic(t *testing.T) {
	assertPanic(t, func() {
		NotNil(nil, "msg")
	})
}

func Test_NotNil_dont_panic(t *testing.T) {
	NotNil(time.Now(), "msg")
}

func Test_NotNilf_panic(t *testing.T) {
	assertPanic(t, func() {
		NotNilf(nil, "msg %v", "msgArg")
	})
}

func Test_NotNilf_dont_panic(t *testing.T) {
	NotNilf(time.Now(), "msg %v", "msgArg")
}

func Test_IsTrue_panic(t *testing.T) {
	assertPanic(t, func() {
		IsTrue(false, "msg")
	})
}

func Test_IsTrue_dont_panic(t *testing.T) {
	IsTrue(true, "msg")
}

func Test_IsTruef_panic(t *testing.T) {
	assertPanic(t, func() {
		IsTruef(false, "msg %v", "msgArg")
	})
}

func Test_IsTruef_dont_panic(t *testing.T) {
	IsTruef(true, "msg %v", "msgArg")
}

func Test_GreaterThan_panic(t *testing.T) {
	assertPanic(t, func() {
		GreaterThan(10, 100, "msg")
	})
}

func Test_GreaterThan_dont_panic(t *testing.T) {
	GreaterThan(100, 10, "msg")
}

func Test_GreaterThanf_panic(t *testing.T) {
	assertPanic(t, func() {
		GreaterThanf(10, 100, "msg %v", "msgArg")
	})
}

func Test_GreaterThanf_dont_panic(t *testing.T) {
	GreaterThanf(100, 10, "msg %v", "msgArg")
}

func Test_LessThan_panic(t *testing.T) {
	assertPanic(t, func() {
		LessThan(100, 10, "msg")
	})
}

func Test_LessThan_dont_panic(t *testing.T) {
	LessThan(10, 100, "msg")
}

func Test_LessThanf_panic(t *testing.T) {
	assertPanic(t, func() {
		LessThanf(100, 10, "msg %v", "msgArg")
	})
}

func Test_LessThanf_dont_panic(t *testing.T) {
	LessThanf(10, 100, "msg %v", "msgArg")
}

func Test_Between_panic(t *testing.T) {
	assertPanic(t, func() {
		Between(1, 10, 100, "msg")
	})
}

func Test_Between_dont_panic(t *testing.T) {
	Between(50, 10, 100, "msg")
}

func Test_Betweenf_panic(t *testing.T) {
	assertPanic(t, func() {
		Betweenf(1, 10, 100, "msg %v", "msgArg")
	})
}

func Test_Betweenf_dont_panic(t *testing.T) {
	Betweenf(50, 10, 100, "msg %v", "msgArg")
}

func Test_Matches_panic(t *testing.T) {
	assertPanic(t, func() {
		Matches("abc", "[d..z]", "msg")
	})
}

func Test_Matches_dont_panic(t *testing.T) {
	Matches("abc", "[a..z]", "msg")
}

func Test_Matchesf_panic(t *testing.T) {
	assertPanic(t, func() {
		Matchesf("abc", "[d..z]", "msg %v", "msgArg")
	})
}

func Test_Matchesf_dont_panic(t *testing.T) {
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
