package rule_validator

type Rule interface {
	Accomplish() bool
}

type Validator struct {
	rules []Rule
}

func NewValidator(r []Rule) Validator {
	return Validator{
		r,
	}
}

func (v Validator) Validate() bool {
	r := false

	for _, rule := range v.rules {
		if rule.Accomplish() {
			r = true
		}
	}

	return r
}
