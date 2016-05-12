// Package argcheck provides functions to perform simple precondition
// checks on arguments.
//
// Usage is really simple, just import package, and right after start
// using any argument in a function, define and apply preconditions.
//
// Example:
//
//  package main
//
//  import (
//  	"fmt"
//  	"argcheck"
//  )
//
//  func calculateMonthlySalary(total int, numMonths int) int {
//  	argcheck.GreaterThanf(numMonths, 0, "numMonths must be positive and not zero, numMonths is %v", numMonths)
//  	return total / numMonths
//  }
//
//  func main() {
//   	total := 50000
//  	numMonths := 0
//  	fmt.Printf("Total salary of %v in %v months, gives %v per month.", total, numMonths, calculateMonthlySalary(total, numMonths))
//  }
//
// Output:
//  panic: numMonths must be positive and not zero, numMonths is [[[0]]]
package argcheck

import (
	"fmt"
	"reflect"
	"regexp"
)

var floatT = reflect.TypeOf(float64(0))

// NotNil checks that argument arg is not nil, otherwise will
// panic showing message msg.
func NotNil(arg interface{}, msg string) {
	NotNilf(arg, msg, nil)
}

// NotNilf checks that argument arg is not nil, otherwise will
// panic showing message msg, formatted with msgArgs.
func NotNilf(arg interface{}, msg string, msgArgs ...interface{}) {
	IsTruef(arg != nil, msg, msgArgs)
}

// IsTrue checks that given condition is true, otherwise will
// panic showing message msg.
func IsTrue(condition bool, msg string) {
	IsTruef(condition, msg, nil)
}

// IsTruef checks that given condition is true, otherwise will
// panic showing message msg, formatted with msgArgs.
func IsTruef(condition bool, msg string, msgArgs ...interface{}) {
	if !condition {
		panic(safeBuildMsg(msg, msgArgs))
	}
}

// GreaterThan checks that argument arg is greater than argument other,
// otherwise will panic showing message msg.
// This comparison is made by converting arg and other to Float64.
func GreaterThan(arg interface{}, other interface{}, msg string) {
	GreaterThanf(arg, other, msg, nil)
}

// GreaterThanf checks that argument arg is greater than argument other,
// otherwise will panic showing message msg, formatted with msgArgs.
// This comparison is made by converting arg and other to Float64.
func GreaterThanf(arg interface{}, other interface{}, msg string, msgArgs ...interface{}) {
	if argAsF, err := convertToFloat64(arg); err != nil {
		panic(err)
	} else if otherAsF, err := convertToFloat64(other); err != nil {
		panic(err)
	} else {
		IsTruef(argAsF > otherAsF, msg, msgArgs)
	}
}

// LessThan checks that argument arg is less than argument other,
// otherwise will panic showing message msg.
// This comparison is made by converting arg and other to Float64.
func LessThan(arg interface{}, other interface{}, msg string) {
	LessThanf(arg, other, msg, nil)
}

// LessThanf checks that argument arg is less than argument other,
// otherwise will panic showing message msg, formatted with msgArgs.
// This comparison is made by converting arg and other to Float64.
func LessThanf(arg interface{}, other interface{}, msg string, msgArgs ...interface{}) {
	if argAsF, err := convertToFloat64(arg); err != nil {
		panic(err)
	} else if otherAsF, err := convertToFloat64(other); err != nil {
		panic(err)
	} else {
		IsTruef(argAsF < otherAsF, msg, msgArgs)
	}
}

// Between checks that argument arg is between arguments from and to,
// otherwise will panic showing message msg.
// This comparison is made by converting arg, from and to to Float64.
func Between(arg interface{}, from interface{}, to interface{}, msg string) {
	Betweenf(arg, from, to, msg, nil)
}

// Between checks that argument arg is between arguments from and to,
// otherwise will panic showing message msg, formatted with msgArgs.
// This comparison is made by converting arg, from and to to Float64.
func Betweenf(arg interface{}, from interface{}, to interface{}, msg string, msgArgs ...interface{}) {
	if argAsF, err := convertToFloat64(arg); err != nil {
		panic(err)
	} else if fromAsF, err := convertToFloat64(from); err != nil {
		panic(err)
	} else if toAsF, err := convertToFloat64(to); err != nil {
		panic(err)
	} else {
		IsTruef(argAsF > fromAsF && argAsF < toAsF, msg, msgArgs)
	}
}

// Matches checks that string argument arg matches the regular expression
// given in r, otherwise will panic showing message msg.
func Matches(arg string, r string, msg string) {
	Matchesf(arg, r, msg, nil)
}

// Matches checks that string argument arg matches the regular expression
// given in r, otherwise will panic showing message msg, formatted with msgArgs.
func Matchesf(arg string, r string, msg string, msgArgs ...interface{}) {
	if matched, err := regexp.MatchString(r, arg); err != nil {
		panic(err)
	} else {
		IsTruef(matched, msg, msgArgs)
	}
}

func safeBuildMsg(msg string, msgArgs ...interface{}) string {
	if msg == "" {
		return ""
	}
	return fmt.Sprintf(msg, msgArgs)
}

func convertToFloat64(n interface{}) (float64, error) {
	v := reflect.ValueOf(n)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatT) {
		return 0, fmt.Errorf("cannot convert %v to float64", v.Type())
	}
	return v.Convert(floatT).Float(), nil
}
