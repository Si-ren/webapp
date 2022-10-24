package error

import "github.com/astaxie/beego/validation"

type Errors struct {
	errors map[string][]string
}

func (e *Errors) AddError(key, err string) {
	if _, ok := e.errors[key]; !ok {
		e.errors[key] = make([]string, 0, 5)
	}
	e.errors[key] = append(e.errors[key], err)
}

func (e *Errors) Errors() map[string][]string {
	return e.errors
}

func (e *Errors) ErrorsByKey(key string) []string {
	return e.errors[key]
}

func (e *Errors) HasErrors() bool {
	return len(e.errors) != 0
}

func New() *Errors {
	return &Errors{
		errors: make(map[string][]string),
	}
}

func (e *Errors) AddValidation(validation *validation.Validation) {
	if validation.HasErrors() {
		for key, errors := range validation.ErrorsMap {
			for _, err := range errors {
				e.AddError(key, err.Message)
			}
		}
	}

}
