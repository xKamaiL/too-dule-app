package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

func IsRFC3339V1(fl validator.FieldLevel) bool {
	fmt.Println("Now: ", fl.Field())
	_, err := time.Parse(time.RFC3339, fl.Field().String())

	if nil != err {
		fmt.Println("check error ", err)
		return false
	}

	return true
}
