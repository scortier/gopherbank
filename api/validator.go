package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/scortier/gopherbank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if currency is valid
		return util.IsSupportCurrency(currency)
	}
	return false
}
