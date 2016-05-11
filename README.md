# go-argcheck
###Simple Go argument preconditions checker.

[![GoDoc](https://godoc.org/github.com/raulbajales/go-argcheck/argcheck?status.svg)](https://godoc.org/github.com/raulbajales/go-argcheck/argcheck)    [![Build Status](https://travis-ci.org/raulbajales/go-argcheck.svg?branch=master)](https://travis-ci.org/raulbajales/go-argcheck)

This minimal library allows to embrace [Fail Fast](https://en.wikipedia.org/wiki/Fail-fast) paradigm, by providing a set of functions to check preconditions and do validations on arguments. Whenever a precondition is not met, a panic will be thrown.