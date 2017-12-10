package simplevalidator_test

import (
	"testing"

	"github.com/jgimeno/simplevalidator"
)

func TestItOnlyNeedsATrueRuleToReturnTrue(t *testing.T) {
	trueRule := TrueRule{}

	v := simplevalidator.NewValidator([]simplevalidator.Rule{trueRule, FalseRule{}})

	r := v.Validate()

	if r == false {
		t.Error("Error using validator.")
	}
}

func TestWithFalseRules(t *testing.T) {
	falseRule := FalseRule{}

	v := simplevalidator.NewValidator([]simplevalidator.Rule{falseRule, falseRule})

	r := v.Validate()

	if r == true {
		t.Error("Error using validator.")
	}
}

type TrueRule struct {
}

func (r TrueRule) Accomplish() bool {
	return true
}

type FalseRule struct {
}

func (r FalseRule) Accomplish() bool {
	return false
}
