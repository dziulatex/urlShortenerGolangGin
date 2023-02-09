package validators

import (
	"github.com/go-playground/validator"
	"goUrlShortener/utils"
	"time"
)

var IsDateGtThanToday validator.Func = func(fl validator.FieldLevel) bool {
	dateTime, ok := fl.Field().Interface().(time.Time)
	date := utils.ParseDateTimeAsOnlyDate(dateTime)
	if ok {
		todayDateTime := time.Now()
		todayDate := utils.ParseDateTimeAsOnlyDate(todayDateTime)
		if todayDate.After(date) || todayDate.Equal(date) {
			return false
		}
	}
	return true
}
var IsDateGtThanTodayOrNull validator.Func = func(fl validator.FieldLevel) bool {
	value := fl.Field().Interface()
	dateTime, ok := value.(time.Time)
	date := utils.ParseDateTimeAsOnlyDate(dateTime)
	if dateTime.IsZero() {
		return true
	}
	if ok {
		todayDateTime := time.Now()
		todayDate := utils.ParseDateTimeAsOnlyDate(todayDateTime)
		if todayDate.After(date) || todayDate.Equal(date) {
			return false
		}
	}
	return true
}
