package validations

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jailtonjunior94/go-tags/enums"
)

func Validations(s any) error {
	v := reflect.ValueOf(s)

	if v.Kind() != reflect.Struct {
		return errors.New("must be a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("validate")

		if tag == "" {
			continue
		}

		val := v.Field(i).String()
		tags := strings.Split(tag, ",")

		for _, t := range tags {
			switch t {
			case "required":
				if val == "" {
					return fmt.Errorf(fmt.Sprintf("%s is required", field.Name))
				}
			case "email":
				if !strings.Contains(val, "@") {
					return fmt.Errorf("%s is not a valid email", field.Name)
				}
			case "person_type":
				if !strings.Contains(val, enums.PessoaFisica.String()) && !strings.Contains(val, enums.PessoaJuridica.String()) {
					return fmt.Errorf("%s is not a valid person type", field.Name)
				}
			}
		}
	}

	return nil
}
