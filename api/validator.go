package api

import (
	"github.com/Al3xDo/simple_bank/util"
	"github.com/go-playground/validator/v10"
)

var validaCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	// .Field -> get value of a field
	// Get value as an empty interface
	// convert value to a string
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported or not
		return util.IsSupportedCurrency(currency)

	}
	return false
}
