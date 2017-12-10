package simplevalidator

// Rule interface that a Rule added to validator should implement.
type Rule interface {
	Accomplish() bool
}

// Validator performs validation based on Rule
type Validator struct {
	rules []Rule
}

// NewValidator Creates a validator.
func NewValidator(r []Rule) Validator {
	return Validator{
		r,
	}
}

// Validate performs a validation and returns bool.
func (v Validator) Validate() bool {
	r := false

	for _, rule := range v.rules {
		if rule.Accomplish() {
			r = true
		}
	}

	return r
}
