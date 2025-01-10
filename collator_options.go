//go:build !darwin
// +build !darwin

package icu4xgo

//#include <icu4xgo.h>
//#include <stdlib.h>
//#include <string.h>
import "C"

type CollatorStrength int

const (
	Primary CollatorStrength = iota
	Secondary
	Tertiary
	Quaternary
	Identical
)

type CollatorAlternateHandling int

const (
	NonIgnorable CollatorAlternateHandling = iota
	Shifted
)

type CollatorCaseFirst int

const (
	Off CollatorCaseFirst = iota
	LowerFirst
	UpperFirst
)

type CollatorMaxVariable int

const (
	Space CollatorMaxVariable = iota
	Punct
	Symbol
	Currency
)

type CollatorCaseLevel int

const (
	CLOff CollatorCaseLevel = iota
	CLOn
)

type CollatorNumeric int

const (
	NumericOff CollatorNumeric = iota
	NumericOn
)

type CollatorBackwardSecondLevel int

const (
	BackwardSecondLevelOff CollatorBackwardSecondLevel = iota
	BackwardSecondLevelOn
)

type CollatorOptions struct {
	Strength            CollatorStrength
	AlternateHandling   CollatorAlternateHandling
	CaseFirst           CollatorCaseFirst
	MaxVariable         CollatorMaxVariable
	CaseLevel           CollatorCaseLevel
	Numeric             CollatorNumeric
	BackwardSecondLevel CollatorBackwardSecondLevel
}

func (c CollatorOptions) ToC() C.IGCollatorOptions {
	return C.IGCollatorOptions{
		strength:              C.CollatorStrength(c.Strength),
		alternate_handling:    C.CollatorAlternateHandling(c.AlternateHandling),
		case_first:            C.CollatorCaseFirst(c.CaseFirst),
		max_variable:          C.CollatorMaxVariable(c.MaxVariable),
		case_level:            C.CollatorCaseLevel(c.CaseLevel),
		numeric:               C.CollatorNumeric(c.Numeric),
		backward_second_level: C.CollatorBackwardSecondLevel(c.BackwardSecondLevel),
	}
}
