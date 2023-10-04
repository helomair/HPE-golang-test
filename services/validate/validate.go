package validate

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func Run(data interface{}, validateType string) []string {
	v := validator.New(validator.WithRequiredStructEnabled())
	var err error

	switch validateType {
	case "struct":
		err = v.Struct(data)
	default:
		err = v.Var(data, validateType)
	}

	ret := verifyError(err)

	return ret
}

func verifyError(err error) []string {
	var ret []string
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
			ret = append(ret, e.Error())
		}
	}

	return ret
}
