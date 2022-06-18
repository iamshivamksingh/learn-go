package reflection

import "reflect"

// any type is alias for interface{}
// walk(x interface{}, fn func(string)) will accept any value for x
// only use reflection if you really need to.

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		// extracing the underlying value
		val = val.Elem()
	}

	return val
}
