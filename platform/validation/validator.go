package validation

type Validator interface {
	Validate(data interface{}) (ok bool, errs []Error)
}

type Error struct {
	FieldName string
	Error     error
}

type ValidatorFunc func(fieldName string, value interface{}, arg string) (bool, error)

func DefaultValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"required": required,
		"min":      min,
	}
}
