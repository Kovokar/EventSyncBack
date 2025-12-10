package validations

import "fmt"

type Field struct {
	Name  string
	Value any
}

func ValidateRequiredFields(fields ...Field) error {
	for _, f := range fields {
		if f.Value == "" {
			return fmt.Errorf("%s é obrigatório", f.Name)
		}
	}
	return nil
}
