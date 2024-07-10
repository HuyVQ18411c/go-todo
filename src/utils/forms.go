package utils

import (
	"reflect"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func AssignValuesToStructFields(data map[string]any, target any) {
	reflectedTarget := reflect.ValueOf(target).Elem()

	for key, value := range data {
		reflectedDataValue := reflect.ValueOf(value)
		structFieldValue := reflectedTarget.FieldByName(cases.Title(language.English).String(key))

		if structFieldValue.Kind() == reflectedDataValue.Kind() {
			structFieldValue.Set(reflectedDataValue)
		} else if reflectedDataValue.CanConvert(structFieldValue.Type()) {
			reflectedDataValue = reflectedDataValue.Convert(structFieldValue.Type())
			reflectedTarget.Set(reflectedDataValue)
		}
	}
}
