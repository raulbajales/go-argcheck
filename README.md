# go-argcheck
### Simple Go argument preconditions checker

[![GoDoc](https://godoc.org/github.com/raulbajales/go-argcheck/argcheck?status.svg)](https://godoc.org/github.com/raulbajales/go-argcheck/argcheck) [![Build Status](https://travis-ci.org/raulbajales/go-argcheck.svg?branch=master)](https://travis-ci.org/raulbajales/go-argcheck)

This minimal library allows to embrace [Fail Fast](https://en.wikipedia.org/wiki/Fail-fast) paradigm, by providing a set of functions to check preconditions and do validations on arguments. Whenever a precondition is not met, a panic will be thrown.

#### Usage
Just get it
```shell
$ go get github.com/raulbajales/go-argcheck/argcheck
```
And use it
```go
import (
	"argcheck"
	"fmt"
)

func calculateMonthlySalary(total int, numMonths int) int {
	argcheck.GreaterThanf(numMonths, 0, "numMonths must be positive and not zero, numMonths is %v", numMonths)
	return total / numMonths
}

func main() {
 	total := 50000
	numMonths := 0
	fmt.Printf("Total salary of %v in %v months, gives %v per month.", total, numMonths, calculateMonthlySalary(total, numMonths))
}
```
Check [GoDoc here](https://godoc.org/github.com/raulbajales/go-argcheck/argcheck) for more preconditions/validations.

###### ~ Made with coffe
