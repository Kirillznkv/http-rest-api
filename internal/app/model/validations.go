package model

import validation "github.com/go-ozzo/ozzo-validation"

func requireadIf(cond bool) validation.RuleFunc {
	return func(val any) error {
		if cond {
			return validation.Validate(val, validation.Required)
		}

		return nil
	}
}
