package main

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

type Validator interface {
	Validate(value interface{}) error
}

type Field struct {
	Name       string
	Value      interface{}
	Validators []Validator
}

type Form struct {
	Fields []Field
}

// ---------------------------------------------
type RequiredValidator struct{}

type LengthValidator struct {
	Min int
	Max int
}

type EmailValidator struct{}

// --------------------------------------------
func (v RequiredValidator) Validate(value interface{}) error {
	if value == nil {
		return errors.New("поле обязательно для заполнения")
	}

	switch val := value.(type) {
	case string:
		if strings.TrimSpace(val) == "" {
			return errors.New("поле обязательно для заполнения")
		}
	case int:
		if val == 0 {
			return errors.New("поле обязательно для заполнения")
		}
	}

	return nil
}

func (v LengthValidator) Validate(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("значение должно быть строкой")
	}

	length := len(strings.TrimSpace(str))

	if v.Min > 0 && length < v.Min {
		return fmt.Errorf("длина должна быть не менее %d символов", v.Min)
	}

	if v.Max > 0 && length > v.Max {
		return fmt.Errorf("длина должна быть не более %d символов", v.Max)
	}

	return nil
}

func (v EmailValidator) Validate(value interface{}) error {
	email, ok := value.(string)
	if !ok {
		return errors.New("email должен быть строкой")
	}

	email = strings.TrimSpace(email)
	if email == "" {
		return nil
	}
	//от regex у меня заболели глаза, поэтому пришлось добавить net/mail в импорт
	//я не уверен что он проверяет правильность написания, а не "существует ли такая почта?", надеюсь что первое
	_, err := mail.ParseAddress(email)
	if !(err == nil) {
		return errors.New("неверный формат email")
	}

	return nil
}

func (f *Form) AddField(name string, value interface{}, validators ...Validator) {
	field := &Field{
		Name:       name,
		Value:      value,
		Validators: validators,
	}
	f.Fields = append(f.Fields, *field)
}

func (f *Form) Validate() bool {
	isValid := true

	for _, field := range f.Fields {
		var fieldErrors []string

		for _, validator := range field.Validators {
			if err := validator.Validate(field.Value); err != nil {
				fieldErrors = append(fieldErrors, err.Error())
			}
		}

		if len(fieldErrors) > 0 {
			//ошибки
			fmt.Printf("ошибки в поле '%s':\n", field.Name)
			for _, err := range fieldErrors {
				fmt.Printf("  - %s\n", err)
			}
			isValid = false
		} else {
			fmt.Printf("поле '%s': OK\n", field.Name)
		}
	}

	return isValid
}

// --------------------------------------------
func main() {
	registrationForm := Form{Fields: nil}

	registrationForm.AddField(
		"username",
		"john_doe",
		RequiredValidator{},
		LengthValidator{Min: 3, Max: 20},
	)

	registrationForm.AddField(
		"email",
		"john@example.com",
		RequiredValidator{},
		EmailValidator{},
	)

	registrationForm.AddField(
		"password",
		"",
		RequiredValidator{},
		LengthValidator{Min: 6, Max: 50},
	)

	fmt.Println("результаты:")
	isValid := registrationForm.Validate()
	fmt.Printf("\nформа валидна: %t\n", isValid)
	fmt.Println("Le fin. 😞")
}
